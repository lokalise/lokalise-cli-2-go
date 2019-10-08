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
	Short: "Keys are core item elements of the project.",
}

var keyListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all keys",
	RunE: func(*cobra.Command, []string) error {
		// preparing filters
		if filterUntranslated {
			keyListOpts.FilterUntranslated = "1"
		}

		resp, err := Api.Keys().WithListOptions(keyListOpts).List(projectId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var keyCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates one or more keys in the project.",
	RunE: func(*cobra.Command, []string) error {
		// preparing options
		err := newKeyFillFields()
		if err != nil {
			return err
		}

		k := Api.Keys()
		// k.SetDebug(false)
		resp, err := k.Create(projectId, []lokalise.NewKey{newKey})
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var keyRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a Key object",
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
	Short: "Updates a Key object",
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
	Short: "Deletes a key from the project.",
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
	fs.Uint8Var(&keyListOpts.DisableReferences, "disable-references", 0, "")
	fs.Uint8Var(&keyListOpts.IncludeComments, "include-comments", 0, "")
	fs.Uint8Var(&keyListOpts.IncludeScreenshots, "include-screenshots", 0, "")
	fs.Uint8Var(&keyListOpts.IncludeTranslations, "include-translations", 0, "")
	fs.StringVar(&keyListOpts.FilterTranslationLangIDs, "filter-translation-lang-ids", "",
		"One or more language ID to filter by (comma separated)")
	fs.StringVar(&keyListOpts.FilterTags, "filter-tags", "", "")
	fs.StringVar(&keyListOpts.FilterFilenames, "filter-filenames", "", "")
	fs.StringVar(&keyListOpts.FilterKeys, "filter-keys", "", "")
	fs.StringVar(&keyListOpts.FilterKeyIDs, "filter-key-ids", "", "")
	fs.StringVar(&keyListOpts.FilterPlatforms, "filter-platforms", "", "Possible values are ios, android, web and other")
	fs.BoolVar(&filterUntranslated, "filter-untranslated", false, "")
	fs.StringVar(&keyListOpts.FilterQAIssues, "filter-qa-issues", "", "")

	// Create
	fs = keyCreateCmd.Flags()
	fs.StringVar(&newKeyName, "key-name", "", "Key identifier")
	_ = keyCreateCmd.MarkFlagRequired("key-name")
	fs.StringVar(&newKey.Description, "description", "", "")
	fs.StringSliceVar(&newKey.Platforms, "platforms", []string{}, "List of platforms, enabled for this key")
	_ = keyCreateCmd.MarkFlagRequired("platforms")
	fs.StringVar(&newKeyFilenames, "filenames", "", "")
	fs.StringSliceVar(&newKey.Tags, "tags", []string{}, "")
	fs.StringSliceVar(&newKeyComments, "comments", []string{}, "")
	// screenshots skipped
	fs.StringVar(&newKeyTranslations, "translations", "", "")
	fs.BoolVar(&newKey.IsPlural, "is-plural", false, "")
	fs.StringVar(&newKey.PluralName, "plural-name", "", "")
	fs.BoolVar(&newKey.IsHidden, "is-hidden", false, "")
	fs.BoolVar(&newKey.IsArchived, "is-archived", false, "")
	fs.StringVar(&newKey.Context, "context", "", "")
	fs.IntVar(&newKey.CharLimit, "char-limit", 0, "")
	fs.StringVar(&newKey.CustomAttributes, "custom-attributes", "", "")

	// Update
	flagKeyId(keyUpdateCmd)
	fs = keyUpdateCmd.Flags()
	fs.StringVar(&newKeyName, "key-name", "", "")
	fs.StringVar(&newKey.Description, "description", "", "")
	fs.StringSliceVar(&newKey.Platforms, "platforms", []string{}, "")
	fs.StringVar(&newKeyFilenames, "filenames", "", "")
	fs.StringSliceVar(&newKey.Tags, "tags", []string{}, "")
	// fs.BoolVar(&newKey.MergeTags, "merge-tags", false, "") // todo enable
	fs.BoolVar(&newKey.IsPlural, "merge-tags", false, "")
	fs.StringVar(&newKey.PluralName, "plural-name", "", "")
	fs.BoolVar(&newKey.IsHidden, "is-hidden", false, "")
	fs.BoolVar(&newKey.IsArchived, "is-archived", false, "")
	fs.StringVar(&newKey.Context, "context", "", "")
	fs.IntVar(&newKey.CharLimit, "char-limit", 0, "")
	fs.StringVar(&newKey.CustomAttributes, "custom-attributes", "", "")

	// retrieve, delete
	flagKeyId(keyRetrieveCmd)
	keyRetrieveCmd.Flags().Uint8Var(&keyListOpts.DisableReferences, "disable-references", 0, "")

	flagKeyId(keyDeleteCmd)
}

func flagKeyId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&keyId, "key-id", 0, "A unique identifier of key (required)")
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
