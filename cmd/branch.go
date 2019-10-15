package cmd

import (
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

var (
	branchId    int64
	branchName string
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Manage branches",
}

var branchListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all branches",
	Long:  "Retrieves a list of project branches. Requires Manage settings admin right.",
	RunE: func(*cobra.Command, []string) error {
		c := Api.Branches()
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

var branchCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a branch",
	Long:  "Creates branch of the project. Requires Manage settings admin right.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Branches().Create(projectId, branchName)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var branchRestoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a branch",
	Long:  "Restores project branch to a project copy. Requires Manage settings admin right and Admin role in the team.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Branches().Restore(projectId, branchId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var branchDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a branch",
	Long:  "Deletes project branch. Requires Manage settings admin right.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Branches().Delete(projectId, branchId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	branchCmd.AddCommand(branchListCmd, branchCreateCmd, branchRestoreCmd, branchDeleteCmd)
	rootCmd.AddCommand(branchCmd)

	// general flags
	flagProjectId(branchCmd, true)

	// separate flags for every command
	branchCreateCmd.Flags().StringVar(&branchName, "name", "", "Branch name.")

	flagBranchId(branchDeleteCmd)
	flagBranchId(branchRestoreCmd)
}

func flagBranchId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&branchId, "branch-id", 0, "A unique identifier of the branch (required).")
	_ = cmd.MarkFlagRequired("branch-id")
}
