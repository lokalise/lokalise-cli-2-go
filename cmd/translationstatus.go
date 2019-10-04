package cmd

import (
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

var (
	trStatusId int64
)

// translationStatusCmd represents the translation-status command
var translationStatusCmd = &cobra.Command{
	Use:   "translationStatus",
	Short: "The ...",
}

var translationStatusListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists project translation statuses",
	RunE: func(cmd *cobra.Command, args []string) error {
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

	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.TranslationStatuses().Create(projectId, lokalise.CreateTranslationStatus{}) // fixme

		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var translationStatusRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a translationStatus ",
	RunE: func(cmd *cobra.Command, args []string) error {
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
	RunE: func(cmd *cobra.Command, args []string) error {
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
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.TranslationStatuses().Update(projectId, trStatusId, lokalise.UpdateTranslationStatus{})
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var translationStatusDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a translationStatus from the project.",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.TranslationStatuses().Delete(projectId, trStatusId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	translationStatusCmd.AddCommand(translationStatusListCmd)
	translationStatusCmd.AddCommand(translationStatusCreateCmd)
	translationStatusCmd.AddCommand(translationStatusRetrieveCmd)
	translationStatusCmd.AddCommand(translationStatusRetrieveColorsCmd)
	translationStatusCmd.AddCommand(translationStatusUpdateCmd)
	translationStatusCmd.AddCommand(translationStatusDeleteCmd)

	rootCmd.AddCommand(translationStatusCmd)

	// general flags
	withProjectId(translationStatusCmd, true)

	// separate flags for every command
	flagTrStatusId(translationStatusRetrieveCmd)
	flagTrStatusId(translationStatusUpdateCmd)
	flagTrStatusId(translationStatusDeleteCmd)
}

func flagTrStatusId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&trStatusId, "status-id", 0, "A unique identifier of translationStatus (required)")
	_ = cmd.MarkFlagRequired("status-id")
}
