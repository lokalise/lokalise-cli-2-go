package cmd

import (
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "The ...",
}

var fileListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists project files",
	/*RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Files().List(projectId, lokalise.FileOptions{})
		if err != nil {
			return err
		}
		return printJson(resp)
	},*/
}

var fileUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Uploads a file to parse",

	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Files().Upload(projectId, lokalise.FileUpload{})

		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var fileDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Downloads a file",

	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Files().Download(projectId, lokalise.FileDownloadOptions{})

		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	fileCmd.AddCommand(fileListCmd)
	fileCmd.AddCommand(fileUploadCmd)
	fileCmd.AddCommand(fileDownloadCmd)

	rootCmd.AddCommand(fileCmd)

	// general flags
	withProjectId(fileCmd, true)

	// separate flags for every command
}
