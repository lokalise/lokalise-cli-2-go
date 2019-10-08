package cmd

import (
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

var (
	trStatusId     int64
	trStatusCreate lokalise.NewTranslationStatus
	trStatusUpdate lokalise.UpdateTranslationStatus
)

// translationStatusCmd represents the translation-status command
var translationStatusCmd = &cobra.Command{
	Use: "translationStatus",
}

var translationStatusListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists project translation statuses",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.TranslationStatuses().List(projectId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var translationStatusCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a translationStatus in the project",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.TranslationStatuses().Create(projectId, trStatusCreate)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var translationStatusRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a translationStatus ",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.TranslationStatuses().Retrieve(projectId, trStatusId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var translationStatusRetrieveColorsCmd = &cobra.Command{
	Use:   "retrieve-colors",
	Short: "Retrieves colors for translation statuses ",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.TranslationStatuses().ListColors(projectId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var translationStatusUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates the properties of a translationStatus",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.TranslationStatuses().Update(projectId, trStatusId, trStatusUpdate)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var translationStatusDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a translationStatus from the project.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.TranslationStatuses().Delete(projectId, trStatusId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	translationStatusCmd.AddCommand(translationStatusListCmd, translationStatusCreateCmd, translationStatusRetrieveCmd,
		translationStatusRetrieveColorsCmd, translationStatusUpdateCmd, translationStatusDeleteCmd)
	rootCmd.AddCommand(translationStatusCmd)

	// general flags
	flagProjectId(translationStatusCmd, true)

	// Create
	fs := translationStatusCreateCmd.Flags()
	fs.StringVar(&trStatusCreate.Title, "title", "", "")
	fs.StringVar(&trStatusCreate.Color, "color", "", "")

	// Update
	flagTrStatusId(translationStatusUpdateCmd)
	fs = translationStatusUpdateCmd.Flags()
	fs.StringVar(&trStatusUpdate.Title, "title", "", "")
	fs.StringVar(&trStatusUpdate.Color, "color", "", "")

	// Retrieve, delete
	flagTrStatusId(translationStatusRetrieveCmd)
	flagTrStatusId(translationStatusDeleteCmd)
}

func flagTrStatusId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&trStatusId, "status-id", 0, "A unique identifier of translationStatus (required)")
	_ = cmd.MarkFlagRequired("status-id")
}
