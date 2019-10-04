package cmd

import (
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "The Project object",
}

var projectListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all projects",
	RunE: func(cmd *cobra.Command, args []string) error {

		resp, err := Api.Projects().List(lokalise.ProjectsOptions{}) // todo alter after changes in the library
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var projectRetrieveCmd = &cobra.Command{
	Use: "retrieve",
	RunE: func(cmd *cobra.Command, args []string) error {

		resp, err := Api.Projects().Retrieve(projectId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var projectUpdateCmd = &cobra.Command{
	Use: "update",
	RunE: func(cmd *cobra.Command, args []string) error {

		resp, err := Api.Projects().Update(projectId, name, description)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var projectDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes the role of a team user. Requires Admin role in the team.",
	RunE: func(cmd *cobra.Command, args []string) error {

		resp, err := Api.Projects().Delete(projectId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var projectEmptyCmd = &cobra.Command{
	Use:   "empty",
	Short: "Deletes the role of a team user. Requires Admin role in the team.",
	RunE: func(cmd *cobra.Command, args []string) error {

		resp, err := Api.Projects().Empty(projectId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	projectCmd.AddCommand(projectListCmd)
	projectCmd.AddCommand(projectRetrieveCmd)
	projectCmd.AddCommand(projectUpdateCmd)
	projectCmd.AddCommand(projectEmptyCmd)
	projectCmd.AddCommand(projectDeleteCmd)

	rootCmd.AddCommand(projectCmd)

}

func withProjectId(cmd *cobra.Command, isPersistent bool) {
	if isPersistent {
		cmd.PersistentFlags().StringVar(&projectId, "project-id", "", "A unique project identifier (required)")
		_ = cmd.MarkPersistentFlagRequired("project-id")
	} else {
		cmd.Flags().StringVar(&projectId, "project-id", "", "A unique project identifier (required)")
		_ = cmd.MarkFlagRequired("project-id")
	}
}
