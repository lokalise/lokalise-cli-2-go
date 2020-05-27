package cmd

import (
	"github.com/lokalise/go-lokalise-api/v3"
	"github.com/spf13/cobra"
)

var (
	trStatusId     int64
	trStatusCreate lokalise.NewTranslationStatus
	trStatusUpdate lokalise.UpdateTranslationStatus
)

// translationStatusCmd represents the translation-status command
var translationStatusCmd = &cobra.Command{
	Use:   "translation-status",
	Short: "Manage custom translation statuses",
	Long:  "Custom translation statuses are used to provide a more efficient translation workflow.",
}

var translationStatusListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all statuses",
	Long:  "Lists all custom translation statuses in the project.",
	RunE: func(*cobra.Command, []string) error {
		c := Api.TranslationStatuses()
		pageOpts := c.PageOpts()

		return repeatableList(
			func(p int64) {
				pageOpts.Page = uint(p)
				c.SetPageOptions(pageOpts)
			},
			func() (lokalise.PageCounter, error) {
				return c.List(projectId)
			},
		)
	},
}

var translationStatusCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a status",
	Long:  "Creates a custom translation status in the project.",
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
	Short: "Retrieve a status",
	Long:  "Retrieves a custom translation status.",
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
	Short: "Retrieve available colors",
	Long:  "Retrieves an array of available colors that can be assigned to custom translation statuses.",
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
	Short: "Update a status",
	Long:  "Updates the custom translation status.",
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
	Short: "Delete a status",
	Long:  "Deletes a custom translation status.",
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
	fs.StringVar(&trStatusCreate.Title, "title", "", "Status title (required).")
	_ = translationStatusCreateCmd.MarkFlagRequired("title")
	fs.StringVar(&trStatusCreate.Color, "color", "", "Hex color of the status (required).")
	_ = translationStatusCreateCmd.MarkFlagRequired("color")

	// Update
	flagTrStatusId(translationStatusUpdateCmd)
	fs = translationStatusUpdateCmd.Flags()
	fs.StringVar(&trStatusUpdate.Title, "title", "", "Status title.")
	fs.StringVar(&trStatusUpdate.Color, "color", "", "Hex color of the status.")

	// Retrieve, delete
	flagTrStatusId(translationStatusRetrieveCmd)
	flagTrStatusId(translationStatusDeleteCmd)
}

func flagTrStatusId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&trStatusId, "status-id", 0, "A unique identifier of the translation status (required).")
	_ = cmd.MarkFlagRequired("status-id")
}
