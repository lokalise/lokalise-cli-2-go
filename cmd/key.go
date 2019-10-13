package cmd

import (
	"encoding/json"
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

var (
	keyId int64

	keyListOpts        lokalise.KeyListOptions
	filterUntranslated bool

	newKey             lokalise.NewKey
	newKeyName         string
	newKeyFilenames    string
	newKeyComments     []string
	newKeyTranslations string
)

// keyCmd represents the key command
var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Manage keys",
	Long: `Keys are core item elements of the project.

Each phrase that is used in your app or website must be identified by a key and have values that represent translations to various languages. For example key index.welcome would have values of Welcome in English and Benvenuto in Italian. Keys can be assigned to one or multiple platforms. Once a key is assigned to a platform, it would be included in the export for file formats related to this platform.
`,
}

var keyListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all keys",
	Long:  "Lists all keys in the project.",
	RunE: func(*cobra.Command, []string) error {
		k := Api.Keys()
		// preparing filters
		keyListOpts.Limit = k.ListOpts().Limit
		if filterUntranslated {
			keyListOpts.FilterUntranslated = "1"
		}

		resp, err := k.WithListOptions(keyListOpts).List(projectId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var keyCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create keys",
	Long:  "Creates one or more keys in the project.",
	RunE: func(*cobra.Command, []string) error {
		// preparing options
		err := newKeyFillFields()
		if err != nil {
			return err
		}

		k := Api.Keys()
		resp, err := k.Create(projectId, []lokalise.NewKey{newKey})
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var keyRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieve a key",
	Long:  "Retrieves a key.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Keys().Retrieve(projectId, keyId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var keyUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a key",
	Long:  "Updates the properties of a key and it’s associated objects. Requires Manage keys admin right.",
	RunE: func(*cobra.Command, []string) error {
		// preparing opts
		err := newKeyFillFields()
		if err != nil {
			return err
		}

		resp, err := Api.Keys().Update(projectId, keyId, newKey)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var keyDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a key",
	Long:  "Deletes a key from the project. Requires Manage keys admin right.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Keys().Delete(projectId, keyId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	keyCmd.AddCommand(keyListCmd, keyCreateCmd, keyRetrieveCmd, keyUpdateCmd, keyDeleteCmd)
	rootCmd.AddCommand(keyCmd)

	// common for all Comment cmd`s
	flagProjectId(keyCmd, true)

	// List
	fs := keyListCmd.Flags()
	fs.Uint8Var(&keyListOpts.DisableReferences, "disable-references", 0, "Whether to disable key references.")
	fs.Uint8Var(&keyListOpts.IncludeComments, "include-comments", 0, "Whether to include comments.")
	fs.Uint8Var(&keyListOpts.IncludeScreenshots, "include-screenshots", 0, "Whether to include URL to screenshots.")
	fs.Uint8Var(&keyListOpts.IncludeTranslations, "include-translations", 0, "Whether to include translations.")
	fs.StringVar(&keyListOpts.FilterTranslationLangIDs, "filter-translation-lang-ids", "",
		"One or more language ID to filter by. Will include translations only for listed IDs.")
	fs.StringVar(&keyListOpts.FilterTags, "filter-tags", "", "One or more tags to filter by.")
	fs.StringVar(&keyListOpts.FilterFilenames, "filter-filenames", "", "One or more filenames to filter by.")
	fs.StringVar(&keyListOpts.FilterKeys, "filter-keys", "", "One or more key name to filter by. In case \"Per-platform keys\" is enabled in project settings, the filter will be applied to all platform names.")
	fs.StringVar(&keyListOpts.FilterKeyIDs, "filter-key-ids", "", "One or more key identifiers to filter by.")
	fs.StringVar(&keyListOpts.FilterPlatforms, "filter-platforms", "", "One or more platforms to filter by. Possible values are ios, android, web and other")
	fs.BoolVar(&filterUntranslated, "filter-untranslated", false, "Filter by untranslated keys.")
	fs.StringVar(&keyListOpts.FilterQAIssues, "filter-qa-issues", "", "One or more QA issues to filter by. Possible values are spelling_and_grammar, placeholders, html, url_count, url, email_count, email, brackets, numbers, leading_whitespace, trailing_whitespace, double_space and special_placeholder.")

	// Create
	fs = keyCreateCmd.Flags()
	fs.StringVar(&newKeyName, "key-name", "", "Key name. For projects with enabled Per-platform key names, pass JSON object with included ios, android, web and other string attributes. (JSON, required).")
	_ = keyCreateCmd.MarkFlagRequired("key-name")
	fs.StringVar(&newKey.Description, "description", "", "Description of the key.")
	fs.StringSliceVar(&newKey.Platforms, "platforms", []string{}, "List of platforms, enabled for this key (required).")
	_ = keyCreateCmd.MarkFlagRequired("platforms")
	fs.StringVar(&newKeyFilenames, "filenames", "", "An object containing key filename attribute for each platform. (JSON, see https://lokalise.com/api2docs/curl/#transition-create-keys-post).")
	fs.StringSliceVar(&newKey.Tags, "tags", []string{}, "List of tags for this keys.")
	fs.StringSliceVar(&newKeyComments, "comments", []string{}, "List of comments for this key.")
	// screenshots skipped
	fs.StringVar(&newKeyTranslations, "translations", "", "Translations for all languages. (JSON, see https://lokalise.com/api2docs/curl/#transition-create-keys-post).")
	fs.BoolVar(&newKey.IsPlural, "is-plural", false, "Whether this key is plural.")
	fs.StringVar(&newKey.PluralName, "plural-name", "", "Optional custom plural name (used in some formats).")
	fs.BoolVar(&newKey.IsHidden, "is-hidden", false, "Whether this key is hidden from non-admins (translators).")
	fs.BoolVar(&newKey.IsArchived, "is-archived", false, "Whether this key is archived.")
	fs.StringVar(&newKey.Context, "context", "", "Optional context of the key (used with some file formats).")
	fs.IntVar(&newKey.CharLimit, "char-limit", 0, "Maximum allowed number of characters in translations for this key.")
	fs.StringVar(&newKey.CustomAttributes, "custom-attributes", "", "JSON containing custom attributes (if any).")

	// Update
	flagKeyId(keyUpdateCmd)
	fs = keyUpdateCmd.Flags()
	fs.StringVar(&newKeyName, "key-name", "", "Key identifier. For projects with enabled Per-platform key names, pass `object` with included ios, android, web and other string attributes.")
	fs.StringVar(&newKey.Description, "description", "", "Description of the key.")
	fs.StringSliceVar(&newKey.Platforms, "platforms", []string{}, "List of platforms, enabled for this key. Possible values are ios, android, web and other.")
	fs.StringVar(&newKeyFilenames, "filenames", "", "An object containing key filename attribute for each platform. (JSON, see https://lokalise.com/api2docs/curl/#transition-update-a-key-put).")
	fs.StringSliceVar(&newKey.Tags, "tags", []string{}, "List of tags for this keys.")
	fs.BoolVar(&newKey.MergeTags, "merge-tags", false, "Enable to merge specified tags with the current tags attached to the key.")
	fs.BoolVar(&newKey.IsPlural, "is-plural", false, "Whether this key is plural.")
	fs.StringVar(&newKey.PluralName, "plural-name", "", "Optional custom plural name (used in some formats).")
	fs.BoolVar(&newKey.IsHidden, "is-hidden", false, "Whether this key is hidden from non-admins (translators).")
	fs.BoolVar(&newKey.IsArchived, "is-archived", false, "Whether this key is archived.")
	fs.StringVar(&newKey.Context, "context", "", "Optional context of the key (used with some file formats).")
	fs.IntVar(&newKey.CharLimit, "char-limit", 0, "Maximum allowed number of characters in translations for this key.")
	fs.StringVar(&newKey.CustomAttributes, "custom-attributes", "", "JSON containing custom attributes (if any).")

	// retrieve, delete
	flagKeyId(keyRetrieveCmd)
	keyRetrieveCmd.Flags().Uint8Var(&keyListOpts.DisableReferences, "disable-references", 0, "Whether to disable key references.")

	flagKeyId(keyDeleteCmd)
}

func flagKeyId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&keyId, "key-id", 0, "A unique identifier of key (required).")
	_ = cmd.MarkFlagRequired("key-id")
}

func newKeyFillFields() error {
	newKey.KeyName = newKeyName

	if newKeyFilenames != "" {
		ps := lokalise.PlatformStrings{}
		err := json.Unmarshal([]byte(newKeyFilenames), &ps)
		if err != nil {
			return err
		}
		newKey.Filenames = &ps
	}

	if newKeyTranslations != "" {
		err := json.Unmarshal([]byte(newKeyTranslations), &newKey.Translations)
		if err != nil {
			return err
		}
	}

	return nil
}
