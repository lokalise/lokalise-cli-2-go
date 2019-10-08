package cmd

import (
	"github.com/lokalise/go-lokalise-api"
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
	Use: "translation",
}

var translationListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists project translations",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Translations().WithListOptions(translationListOpts).List(projectId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var translationRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a translation ",
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
	Short: "Updates a translation from the project.",
	RunE: func(*cobra.Command, []string) error {
		// processing opts
		translationUpdate.IsFuzzy = &translationUpdateIsFuzzy

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
	fs.Uint8Var(&translationListOpts.DisableReferences, "disable-references", 0, "")
	fs.StringVar(&translationListOpts.FilterLangID, "filter-lang-id", "", "")
	fs.Uint8Var(&translationListOpts.FilterIsReviewed, "filter-is-reviewed", 0, "")
	fs.Uint8Var(&translationListOpts.FilterFuzzy, "filter-fuzzy", 0, "")
	fs.StringVar(&translationListOpts.FilterQAIssues, "filter-qa-issues", "", "")

	// Retrieve
	flagTranslationId(translationRetrieveCmd)

	// Update
	flagTranslationId(translationUpdateCmd)
	fs = translationUpdateCmd.Flags()
	fs.StringVar(&translationUpdate.Translation, "translation", "", "")
	_ = translationUpdateCmd.MarkFlagRequired("translation")
	fs.BoolVar(&translationUpdateIsFuzzy, "is-fuzzy", true, "")
	fs.BoolVar(&translationUpdate.IsReviewed, "is-reviewed", false, "")
}

func flagTranslationId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&translationId, "translation-id", 0, "A unique identifier of translation (required)")
	_ = cmd.MarkFlagRequired("translation-id")
}
