package cmd

import (
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

var (
	translationId int64
)

// translationCmd represents the translation command
var translationCmd = &cobra.Command{
	Use:   "translation",
	Short: "The ...",
}

var translationListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists project translations",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Translations().List(projectId, lokalise.TranslationsOptions{})
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var translationRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a translation ",
	RunE: func(cmd *cobra.Command, args []string) error {
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
	/*RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Translations().Update(projectId, translationId) // todo combine parameters
		if err != nil {
			return err
		}
		return printJson(resp)
	},*/
}

func init() {
	translationCmd.AddCommand(translationListCmd)
	translationCmd.AddCommand(translationRetrieveCmd)
	translationCmd.AddCommand(translationUpdateCmd)

	rootCmd.AddCommand(translationCmd)

	// general flags
	withProjectId(translationCmd, true)

	// separate flags for every command
	flagTranslationId(translationRetrieveCmd)
	flagTranslationId(translationUpdateCmd)
}

func flagTranslationId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&translationId, "translation-id", 0, "A unique identifier of translation (required)")
	_ = cmd.MarkFlagRequired("translation-id")
}
