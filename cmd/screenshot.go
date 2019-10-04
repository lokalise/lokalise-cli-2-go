package cmd

import (
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

var (
	screenshotId int64
)

// screenshotCmd represents the screenshot command
var screenshotCmd = &cobra.Command{
	Use:   "screenshot",
	Short: "The ...",
}

var screenshotListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists project screenshots",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Screenshots().List(lokalise.ScreenshotsOptions{}) // todo wtf?! should be only project-id
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var screenshotCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a screenshot in the project",

	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Screenshots().Create(projectId, lokalise.CreateScreenshotOptions{}) // fixme

		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var screenshotRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a screenshot ",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Screenshots().Retrieve(projectId, screenshotId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var screenshotUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates the properties of a screenshot",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Screenshots().Update(projectId, screenshotId, lokalise.UpdateScreenshotOptions{})
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var screenshotDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a screenshot from the project.",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Screenshots().Delete(projectId, screenshotId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	screenshotCmd.AddCommand(screenshotListCmd)
	screenshotCmd.AddCommand(screenshotCreateCmd)
	screenshotCmd.AddCommand(screenshotRetrieveCmd)
	screenshotCmd.AddCommand(screenshotUpdateCmd)
	screenshotCmd.AddCommand(screenshotDeleteCmd)

	rootCmd.AddCommand(screenshotCmd)

	// general flags
	withProjectId(screenshotCmd, true)

	// separate flags for every command
	flagScreenshotId(screenshotCreateCmd)
	flagScreenshotId(screenshotRetrieveCmd)
	flagScreenshotId(screenshotDeleteCmd)
}

func flagScreenshotId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&screenshotId, "screenshot-id", 0, "A unique identifier of screenshot (required)")
	_ = cmd.MarkFlagRequired("screenshot-id")
}
