package cmd

import (
	"github.com/lokalise/go-lokalise-api/v3"
	"github.com/spf13/cobra"
)

var (
	translationId int64

	translationListOpts      lokalise.TranslationListOptions
	translationUpdate        lokalise.UpdateTranslation
	translationUpdateIsFuzzy bool
)

// translationCmd represents the translation command
var translationCmd = &cobra.Command{
	Use:   "translation",
	Short: "Manage translations",
}

var translationListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all translations",
	Long:  "Retrieves a list of project translation items, ungrouped. You may want to request Keys resource in order to get the structured key/translation pairs for all languages.",
	RunE: func(*cobra.Command, []string) error {
		t := Api.Translations()
		translationListOpts.Limit = t.ListOpts().Limit

		return repeatableList(
			func(p int64) {
				translationListOpts.Page = uint(p)
				t.SetListOptions(translationListOpts)
			},
			func() (lokalise.PageCounter, error) {
				return t.List(projectId)
			},
		)
	},
}

var translationRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieve a translation ",
	Long:  "Retrieves a translation.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Translations().Retrieve(projectId, translationId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var translationUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a translation",
	Long:  "Updates a translation.",
	RunE: func(*cobra.Command, []string) error {
		// processing opts
		translationUpdate.IsUnverified = &translationUpdateIsFuzzy

		resp, err := Api.Translations().Update(projectId, translationId, translationUpdate)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	translationCmd.AddCommand(translationListCmd, translationRetrieveCmd, translationUpdateCmd)
	rootCmd.AddCommand(translationCmd)

	// general flags
	flagProjectId(translationCmd, true)

	// List
	fs := translationListCmd.Flags()
	fs.Uint8Var(&translationListOpts.DisableReferences, "disable-references", 0, "Whether to disable key references.")
	fs.Int64Var(&translationListOpts.FilterLangID, "filter-lang-id", 0, "Return translations only for presented language ID.")
	fs.Uint8Var(&translationListOpts.FilterIsReviewed, "filter-is-reviewed", 0, "Filter translations which are reviewed.")
	fs.Uint8Var(&translationListOpts.FilterUnverified, "filter-fuzzy", 0, "Filter translations which are unverified (fuzzy).")
	fs.StringVar(&translationListOpts.FilterQAIssues, "filter-qa-issues", "", "One or more QA issues to filter by. Possible values are spelling_and_grammar, placeholders, html, url_count, url, email_count, email, brackets, numbers, leading_whitespace, trailing_whitespace, double_space and special_placeholder.")

	// Retrieve
	flagTranslationId(translationRetrieveCmd)

	// Update
	flagTranslationId(translationUpdateCmd)
	fs = translationUpdateCmd.Flags()
	fs.StringVar(&translationUpdate.Translation, "translation", "", "The actual translation content. Use a JSON object for plural keys (required).")
	_ = translationUpdateCmd.MarkFlagRequired("translation")
	fs.BoolVar(&translationUpdateIsFuzzy, "is-fuzzy", false, "Whether the Fuzzy flag is enabled. (Note: Fuzzy is called Unverified in the editor now) .")
	fs.BoolVar(&translationUpdate.IsReviewed, "is-reviewed", false, "Whether the Reviewed flag is enabled.")
	fs.StringSliceVar(&translationUpdate.CustomTranslationStatusIDs, "custom-translation-status-ids", []string{}, "Custom translation status IDs to assign to translation (existing statuses will be replaced).")
}

func flagTranslationId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&translationId, "translation-id", 0, "A unique identifier of the translation (required).")
	_ = cmd.MarkFlagRequired("translation-id")
}
