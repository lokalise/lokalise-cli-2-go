package cmd

import (
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

var (
	downloadOpts              lokalise.FileDownload
	downloadOptsReplaceBreaks bool
	downloadOptsLangMapping   string

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
	Use: "file",
}

var fileListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists project files",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Files().List(projectId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var fileUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Uploads a file to parse",
	RunE: func(*cobra.Command, []string) error {
		// preparing opts
		uploadOpts.ConvertPlaceholders = &uploadOptsConvertPlaceholders
		uploadOpts.TagInsertedKeys = &uploadOptsTagInsertedKeys
		uploadOpts.TagUpdatedKeys = &uploadOptsTagUpdatedKeys

		f := Api.Files()
		f.SetDebug(true)
		resp, err := f.Upload(projectId, uploadOpts)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var fileDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Downloads a file",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Files().Download(projectId, downloadOpts)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	fileCmd.AddCommand(fileListCmd, fileUploadCmd, fileDownloadCmd)
	rootCmd.AddCommand(fileCmd)

	// general flags
	flagProjectId(fileCmd, true)

	// Download
	fs := fileDownloadCmd.Flags()
	fs.StringVar(&downloadOpts.Format, "format", "", "File format, e.g. json, strings, xml. (required)")
	_ = fileDownloadCmd.MarkFlagRequired("format")

	fs.BoolVar(&downloadJsonOnly, "json-only", false, "Should only the API JSON response be returned")
	fs.BoolVar(&downloadKeepZip, "keep-zip", false, "keep or delete zip after being unpacked")
	fs.StringVar(&downloadDestination, "dest", "/tmp", "destination folder for zip file")
	fs.StringVar(&downloadUnzipTo, "unzip-to", "", "unzip to this folder")

	fs.BoolVar(&downloadOpts.OriginalFilenames, "original-filenames", true, "")
	// fs.StringVar(&downloadOpts.BundleStructure, "bundle-structure", "", "")
	fs.StringVar(&downloadOpts.DirectoryPrefix, "directory-prefix", "", "")
	fs.BoolVar(&downloadOpts.AllPlatforms, "all-platforms", false, "")
	fs.StringVar(&downloadOpts.FilterLangs, "filter-langs", "", "")
	fs.StringVar(&downloadOpts.FilterData, "filter-data", "", "")
	fs.StringVar(&downloadOpts.FilterFilenames, "filter-filenames", "", "")
	fs.BoolVar(&downloadOpts.AddNewlineEOF, "add-newline-eof", false, "")
	fs.StringVar(&downloadOpts.CustomTranslationStatusIDs, "custom-translation-status-ids", "", "")
	fs.StringVar(&downloadOpts.IncludeTags, "include-tags", "", "")
	fs.StringVar(&downloadOpts.ExcludeTags, "exclude-tags", "", "")
	fs.StringVar(&downloadOpts.ExportSort, "export-sort", "", "")
	fs.StringVar(&downloadOpts.ExportEmptyAs, "export-empty-as", "", "")
	fs.BoolVar(&downloadOpts.IncludeComments, "include-comments", false, "")
	fs.BoolVar(&downloadOpts.IncludeDescription, "include-description", false, "")
	fs.StringVar(&downloadOpts.IncludeProjectIDs, "include-pids", "", "")
	fs.StringVar(&downloadOpts.Triggers, "triggers", "", "")
	fs.StringVar(&downloadOpts.FilterRepositories, "filter-repositories", "", "")
	fs.BoolVar(&downloadOptsReplaceBreaks, "replace-breaks", true, "")
	fs.BoolVar(&downloadOpts.DisableReferences, "disable-references", false, "")
	fs.StringVar(&downloadOpts.PluralFormat, "plural-format", "", "")
	fs.StringVar(&downloadOpts.PlaceholderFormat, "placeholder-format", "", "")
	fs.StringVar(&downloadOpts.WebhookURL, "webhook-url", "", "")
	fs.StringVar(&downloadOptsLangMapping, "language-mapping", "", "")
	fs.BoolVar(&downloadOpts.ICUNumeric, "icu-numeric", false, "")
	fs.BoolVar(&downloadOpts.EscapePercent, "escape-percent", false, "")
	fs.StringVar(&downloadOpts.Indentation, "indentation", "", "")
	fs.BoolVar(&downloadOpts.YAMLIncludeRoot, "yaml-include-root", false, "")
	fs.BoolVar(&downloadOpts.JSONUnescapedSlashes, "json-unescaped-slashes", false, "")
	fs.StringVar(&downloadOpts.JavaPropertiesEncoding, "java-properties-encoding", "", "")
	fs.StringVar(&downloadOpts.JavaPropertiesSeparator, "java-properties-separator", "", "")
	fs.StringVar(&downloadOpts.BundleDescription, "bundle-description", "", "")

	// Upload
	fs = fileUploadCmd.Flags()
	fs.StringVar(&uploadFile, "file", "", "Path to file")
	_ = fileUploadCmd.MarkFlagRequired("file")
	// force-filename is skipped because current time only single-file is supplied
	fs.StringVar(&uploadOpts.LangISO, "lang-iso", "", "")
	_ = fileUploadCmd.MarkFlagRequired("lang-iso")
	fs.BoolVar(&uploadOptsConvertPlaceholders, "convert-placeholders", true, "")
	fs.BoolVar(&uploadOpts.DetectICUPlurals, "detect-icu-plurals", false, "")
	fs.StringSliceVar(&uploadOpts.Tags, "tags", []string{}, "")
	fs.BoolVar(&uploadOptsTagInsertedKeys, "tag-inserted-keys", true, "")
	fs.BoolVar(&uploadOptsTagUpdatedKeys, "tag-updated-keys", true, "")
	fs.BoolVar(&uploadOpts.TagSkippedKeys, "tag-skipped-keys", false, "")
	fs.BoolVar(&uploadOpts.ReplaceModified, "replace-modified", false, "")
	fs.BoolVar(&uploadOpts.SlashNToLinebreak, "slashn-to-linebreak", false, "")
	fs.BoolVar(&uploadOpts.KeysToValues, "keys-to-values", false, "")
	fs.BoolVar(&uploadOpts.DistinguishByFile, "distinguish-by-file", false, "")
	fs.BoolVar(&uploadOpts.ApplyTM, "apply-tm", false, "")
	fs.BoolVar(&uploadOpts.CleanupMode, "cleanup-mode", false, "")
}
