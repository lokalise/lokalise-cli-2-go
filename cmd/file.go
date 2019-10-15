package cmd

import (
	"archive/zip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

var (
	filterFilename string

	downloadOpts                  lokalise.FileDownload
	downloadOptsReplaceBreaks     bool
	downloadOptsOriginalFilenames bool
	downloadOptsLangMapping       string

	downloadJsonOnly    bool
	downloadDestination string
	downloadUnzipTo     string
	downloadKeepZip     bool

	uploadOpts                    lokalise.FileUpload
	uploadOptsConvertPlaceholders bool
	uploadOptsTagInsertedKeys     bool
	uploadOptsTagUpdatedKeys      bool

	uploadFile string
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Upload and download files",
	Long:  "Lokalise is a project-oriented translation management system, which means we store all keys and translations in the database and can generate files in any format you require. Assigning a key to one or more platforms means including the key in the export routine for file formats, associated with this platform, e.g. if a key is assigned to iOS platform it would get included with strings and xliff format exports. In addition to assign keys to platforms you may assign keys to files and have different filename depending on the platform. List of supported file formats is available here https://docs.lokalise.com/en/collections/652248-supported-file-formats.",
}

var fileListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all files",
	Long:  "Lists project files and associated key count. If there are some keys in the project that do not have a file association, they will be returned with filename __unassigned__.",
	RunE: func(*cobra.Command, []string) error {
		c := Api.Files()
		listOpts := c.ListOpts()
		listOpts.Filename = filterFilename

		return repeatableList(
			func(p int64) {
				listOpts.Page = uint(p)
				c.SetListOptions(listOpts)
			},
			func() (lokalise.PageCounter, error) {
				return c.List(projectId)
			},
		)
	},
}

var fileUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a file",
	Long:  "Imports a localization file to the project. Requires Upload files admin right. List of supported file formats is available here https://docs.lokalise.com/en/collections/652248-supported-file-formats",
	RunE: func(*cobra.Command, []string) error {
		f := Api.Files()

		// preparing opts
		uploadOpts.ConvertPlaceholders = &uploadOptsConvertPlaceholders
		uploadOpts.TagInsertedKeys = &uploadOptsTagInsertedKeys
		uploadOpts.TagUpdatedKeys = &uploadOptsTagUpdatedKeys

		files, err := filepath.Glob(uploadFile)
		if err != nil {
			return err
		}

		for _, file := range files {
			fmt.Println("Uploading", file+"...")
			buf, err := ioutil.ReadFile(file)
			if err != nil {
				return err
			}

			uploadOpts.Data = base64.StdEncoding.EncodeToString(buf)
			uploadOpts.Filename = path.Base(file)

			resp, err := f.Upload(projectId, uploadOpts)
			if err != nil {
				return err
			}

			_ = printJson(resp)
		}

		return nil
	},
}

var fileDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download files",
	Long:  "Exports project files as a .zip bundle. Generated bundle will be uploaded to an Amazon S3 bucket, which will be stored there for 12 months available to download. As the bundle is generated and uploaded you would get a response with the URL to the file. Requires Download files admin right.",
	RunE: func(*cobra.Command, []string) error {
		// preparing options
		if downloadOptsLangMapping != "" {
			var mappings []lokalise.LanguageMapping
			err := json.Unmarshal([]byte(downloadOptsLangMapping), &mappings)
			if err != nil {
				return err
			}
			downloadOpts.LanguageMapping = mappings
		}

		downloadOpts.ReplaceBreaks = &downloadOptsReplaceBreaks
		downloadOpts.OriginalFilenames = &downloadOptsOriginalFilenames

		if !downloadJsonOnly {
			fmt.Print("Requesting... ")
		}

		resp, err := Api.Files().Download(projectId, downloadOpts)
		if err != nil {
			return err
		}

		if !downloadJsonOnly {
			fmt.Println("OK")
		}

		if downloadJsonOnly {
			return printJson(resp)
		} else {
			fmt.Println("Downloading", resp.BundleURL+"...")
		}

		err = downloadAndUnzip(resp.BundleURL, downloadDestination, downloadUnzipTo)

		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	fileCmd.AddCommand(fileListCmd, fileUploadCmd, fileDownloadCmd)
	rootCmd.AddCommand(fileCmd)

	// general flags
	flagProjectId(fileCmd, true)

	// List
	fileListCmd.Flags().StringVar(&filterFilename, "filter-filename", "", "Set filename filter for the list.")

	// Download
	fs := fileDownloadCmd.Flags()
	fs.StringVar(&downloadOpts.Format, "format", "", "File format (e.g. json, strings, xml). Must be file extension of any of the file formats we support. May also be ios_sdk or android_sdk for respective OTA SDK bundles. (required)")
	_ = fileDownloadCmd.MarkFlagRequired("format")

	fs.BoolVar(&downloadJsonOnly, "json-only", false, "Should only the API JSON response be returned.")
	fs.BoolVar(&downloadKeepZip, "keep-zip", false, "Keep or delete ZIP file after being unpacked.")
	fs.StringVar(&downloadDestination, "dest", "./", "Destination folder for ZIP file.")
	fs.StringVar(&downloadUnzipTo, "unzip-to", "./", "Unzip to this folder.")

	fs.BoolVar(&downloadOptsOriginalFilenames, "original-filenames", true, "Enable to use original filenames/formats. If set to false all keys will be export to a single file per language (default true).")
	fs.StringVar(&downloadOpts.BundleStructure, "bundle-structure", "", "Bundle structure, used when original-filenames set to false. Allowed placeholders are %LANG_ISO%, %LANG_NAME%, %FORMAT% and %PROJECT_NAME%).")
	fs.StringVar(&downloadOpts.DirectoryPrefix, "directory-prefix", "", "Directory prefix in the bundle, used when original_filenames set to true). Allowed placeholder is %LANG_ISO%.")
	fs.BoolVar(&downloadOpts.AllPlatforms, "all-platforms", false, "Enable to include all platform keys. If disabled, only the keys, associated with the platform of the format will be exported.")
	fs.StringSliceVar(&downloadOpts.FilterLangs, "filter-langs", []string{}, "List of languages to export. Omit this parameter for all languages.")
	fs.StringSliceVar(&downloadOpts.FilterData, "filter-data", []string{}, "Narrow export data range. Allowed values are translated or untranslated, reviewed (or reviewed_only), last_reviewed_only, nonfuzzy and nonhidden. (Note: Fuzzy is called Unverified in the editor now).")
	fs.StringSliceVar(&downloadOpts.FilterFilenames, "filter-filenames", []string{}, "Only keys attributed to selected files will be included. Leave empty for all.")
	fs.BoolVar(&downloadOpts.AddNewlineEOF, "add-newline-eof", false, "Enable to add new line at end of file (if supported by format).")
	fs.StringSliceVar(&downloadOpts.CustomTranslationStatusIDs, "custom-translation-status-ids", []string{}, "Only translations attributed to selected custom statuses will be included. Leave empty for all.")
	fs.StringSliceVar(&downloadOpts.IncludeTags, "include-tags", []string{}, "Narrow export range to tags specified.")
	fs.StringSliceVar(&downloadOpts.ExcludeTags, "exclude-tags", []string{}, "Specify to exclude keys with these tags.")
	fs.StringVar(&downloadOpts.ExportSort, "export-sort", "", "Export key sort mode. Allowed value are first_added, last_added, last_updated, a_z, z_a.")
	fs.StringVar(&downloadOpts.ExportEmptyAs, "export-empty-as", "", "Select how you would like empty translations to be exported. Allowed values are empty to keep empty, base to replace with the base language value, or skip to omit.")
	fs.BoolVar(&downloadOpts.IncludeComments, "include-comments", false, "Enable to include key comments and description in exported file (if supported by the format).")
	fs.BoolVar(&downloadOpts.IncludeDescription, "include-description", false, "Enable to include key description in exported file (if supported by the format).")
	fs.StringSliceVar(&downloadOpts.IncludeProjectIDs, "include-pids", []string{}, "Other projects ID's, which keys should be included with this export.")
	fs.StringSliceVar(&downloadOpts.Triggers, "triggers", []string{}, "Trigger integration exports (must be enabled in project settings). Allowed values are amazons3, gcs, github, gitlab, bitbucket.")
	fs.StringSliceVar(&downloadOpts.FilterRepositories, "filter-repositories", []string{}, "Pull requests will be created only for listed repositories (organization/repository format). Leave empty array to process all configured integrations by platform only.")
	fs.BoolVar(&downloadOptsReplaceBreaks, "replace-breaks", true, "Enable to replace line breaks in exported translations with \\n (default true).")
	fs.BoolVar(&downloadOpts.DisableReferences, "disable-references", false, "Enable to skip automatic replace of key reference placeholders (e.g. [%key:hello_world%]) with their corresponding translations.")
	fs.StringVar(&downloadOpts.PluralFormat, "plural-format", "", "Override the default plural format for the file type. Allowed values are json_string, icu, array, generic, symfony.")
	fs.StringVar(&downloadOpts.PlaceholderFormat, "placeholder-format", "", "Override the default placeholder format for the file type. Allowed values are printf, ios, icu, net, symfony.")
	fs.StringVar(&downloadOpts.WebhookURL, "webhook-url", "", "Once the export is complete, sends a HTTP POST with the generated bundle URL to the specified URL.")
	fs.StringVar(&downloadOptsLangMapping, "language-mapping", "", "List of languages to override default iso codes for this export (JSON, see https://lokalise.com/api2docs/curl/#transition-download-files-post).")
	fs.BoolVar(&downloadOpts.ICUNumeric, "icu-numeric", false, "If enabled, plural forms zero, one and two will be replaced with =0, =1 and =2 respectively. Only works for ICU plural format.")
	fs.BoolVar(&downloadOpts.EscapePercent, "escape-percent", false, "Only works for printf placeholder format. When enabled, all universal percent placeholders \"[%]\" will be always exported as \"%%\".")
	fs.StringVar(&downloadOpts.Indentation, "indentation", "", "Provide to override default indentation in supported files. Allowed values are default, 1sp, 2sp, 3sp, 4sp, 5sp, 6sp, 7sp, 8sp and tab.")
	fs.BoolVar(&downloadOpts.YAMLIncludeRoot, "yaml-include-root", false, " (YAML export only). Enable to include language ISO code as root key.")
	fs.BoolVar(&downloadOpts.JSONUnescapedSlashes, "json-unescaped-slashes", false, "(JSON export only). Enable to leave forward slashes unescaped.")
	fs.StringVar(&downloadOpts.JavaPropertiesEncoding, "java-properties-encoding", "", "(Java Properties export only). Encoding for .properties files. Allowed values are utf-8 and latin-1.")
	fs.StringVar(&downloadOpts.JavaPropertiesSeparator, "java-properties-separator", "", "(Java Properties export only). Separator for keys/values in .properties files. Allowed values are = and :.")
	fs.StringVar(&downloadOpts.BundleDescription, "bundle-description", "", "Description of the created bundle. Applies to ios_sdk or android_sdk OTA SDK bundles.")

	// Upload
	fs = fileUploadCmd.Flags()
	fs.StringVar(&uploadFile, "file", "", "Path to local file (required).")
	_ = fileUploadCmd.MarkFlagRequired("file")
	// force-filename is skipped because current time only single-file is supplied
	fs.StringVar(&uploadOpts.LangISO, "lang-iso", "", "Language code of the translations in the file you are importing (required).")
	_ = fileUploadCmd.MarkFlagRequired("lang-iso")
	fs.BoolVar(&uploadOptsConvertPlaceholders, "convert-placeholders", false, "Enable to automatically convert placeholders to the Lokalise universal placeholders.")
	fs.BoolVar(&uploadOpts.DetectICUPlurals, "detect-icu-plurals", false, "Enable to automatically detect and parse ICU formatted plurals in your translations.")
	fs.StringSliceVar(&uploadOpts.Tags, "tags", []string{}, "Tag keys with the specified tags. By default tags are applied to created and updated keys.")
	fs.BoolVar(&uploadOptsTagInsertedKeys, "tag-inserted-keys", true, "Add specified tags to inserted keys (default true).")
	fs.BoolVar(&uploadOptsTagUpdatedKeys, "tag-updated-keys", true, "Add specified tags to updated keys (default true).")
	fs.BoolVar(&uploadOpts.TagSkippedKeys, "tag-skipped-keys", false, "Add specified tags to skipped keys.")
	fs.BoolVar(&uploadOpts.ReplaceModified, "replace-modified", false, "Enable to replace translations, that have been modified (in the file being uploaded).")
	fs.BoolVar(&uploadOpts.SlashNToLinebreak, "slashn-to-linebreak", false, "Enable to replace \\n with a line break.")
	fs.BoolVar(&uploadOpts.KeysToValues, "keys-to-values", false, "Enable to automatically replace values with key names.")
	fs.BoolVar(&uploadOpts.DistinguishByFile, "distinguish-by-file", false, "Enable to allow keys with similar names to coexist, in case they are assigned to differrent filenames.")
	fs.BoolVar(&uploadOpts.ApplyTM, "apply-tm", false, "Enable to automatically apply 100% translation memory matches.")
	fs.BoolVar(&uploadOpts.CleanupMode, "cleanup-mode", false, "Enable to delete all keys with all language translations that are not present in the uploaded file. You may want to make a snapshot of the project before importing new file, just in case.")
}

//noinspection GoUnhandledErrorResult
func downloadAndUnzip(srcUrl, destPath, unzipPath string) error {
	fileName := path.Base(srcUrl)
	zip, err := os.Create(path.Join(destPath, fileName))
	if err != nil {
		return err
	}
	defer zip.Close()

	resp, err := http.Get(srcUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(zip, resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Unzipping to", unzipPath+"...")
	err = unzip(zip.Name(), unzipPath)
	if err != nil {
		return err
	}

	if !downloadKeepZip {
		_ = os.Remove(zip.Name())
	}

	return nil
}

//noinspection GoUnhandledErrorResult
func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
