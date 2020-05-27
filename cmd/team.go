package cmd

import (
	"github.com/lokalise/go-lokalise-api/v3"
	"github.com/spf13/cobra"
)

var (
	teamId int64
)

// teamCmd represents the team command
var teamCmd = &cobra.Command{
	Use:   "team",
	Short: "List teams",
}

// teamListCmd represents team list command
var teamListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all teams",
	Long:  "Lists all teams available to the user.",
	RunE: func(*cobra.Command, []string) error {
		c := Api.Teams()
		pageOpts := c.PageOpts()

		return repeatableList(
			func(p int64) {
				pageOpts.Page = uint(p)
				c.SetPageOptions(pageOpts)
			},
			func() (lokalise.PageCounter, error) {
				return c.List()
			},
		)
	},
}

func init() {
	teamCmd.AddCommand(teamListCmd)
	rootCmd.AddCommand(teamCmd)
}

// always persistent
func flagTeamId(cmd *cobra.Command) {
	cmd.PersistentFlags().Int64Var(&teamId, "team-id", 0, "A unique identifier of the team (required).")
	_ = cmd.MarkPersistentFlagRequired("team-id")
}
