package cmd

import (
	"github.com/spf13/cobra"
)

// teamCmd represents the team command
var teamCmd = &cobra.Command{
	Use:   "team",
	Short: "The Team object",
}

// teamListCmd represents team list command
var teamListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all teams available to the user",
	RunE: func(cmd *cobra.Command, args []string) error {

		resp, err := Api.Teams().List()
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	teamCmd.AddCommand(teamListCmd)

	rootCmd.AddCommand(teamCmd)
}
