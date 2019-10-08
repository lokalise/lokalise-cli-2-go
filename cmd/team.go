package cmd

import (
	"github.com/spf13/cobra"
)

var (
	teamId int64
)

// teamCmd represents the team command
var teamCmd = &cobra.Command{
	Use:   "team",
}

// teamListCmd represents team list command
var teamListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all teams available to the user",
	RunE: func(*cobra.Command, []string) error {

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

// always persistent
func flagTeamId(cmd *cobra.Command) {
	cmd.PersistentFlags().Int64Var(&teamId, "team-id", 0, "A unique identifier of team (required)")
	_ = cmd.MarkPersistentFlagRequired("team-id")
}
