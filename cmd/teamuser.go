package cmd

import (
	"github.com/lokalise/go-lokalise-api/v4"
	"github.com/spf13/cobra"
)

var (
	userId int64
	role   string
)

// teamUserCmd represents the team-user command
var teamUserCmd = &cobra.Command{
	Use:   "team-user",
	Short: "Manage team users",
}

// teamUserListCmd represents team-user list command
var teamUserListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all team users",
	Long:  "List all team users. Requires Admin role in the team.",
	RunE: func(*cobra.Command, []string) error {
		c := Api.TeamUsers()
		pageOpts := c.PageOpts()

		return repeatableList(
			func(p int64) {
				pageOpts.Page = uint(p)
				c.SetPageOptions(pageOpts)
			},
			func() (lokalise.PageCounter, error) {
				return c.List(teamId)
			},
		)
	},
}

// teamUserRetrieveCmd represents team-user retrieve command
var teamUserRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a team user",
	Long:  "Retrieves a team user. Requires Admin role in the team.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.TeamUsers().Retrieve(teamId, userId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

// teamUserUpdateCmd represents team-user update command
var teamUserUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a team user",
	Long:  "Updates the role of a team user. Requires Admin role in the team.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.TeamUsers().UpdateRole(teamId, userId, lokalise.TeamUserRole(role))
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

// teamUserDeleteCmd represents team-user delete command
var teamUserDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a team user",
	Long:  "Deletes a user from the team. Requires Admin role in the team.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.TeamUsers().Delete(teamId, userId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	teamUserCmd.AddCommand(teamUserListCmd, teamUserRetrieveCmd, teamUserUpdateCmd, teamUserDeleteCmd)
	rootCmd.AddCommand(teamUserCmd)

	// General
	flagTeamId(teamUserCmd)

	// Update
	flagUserId(teamUserUpdateCmd)
	teamUserUpdateCmd.Flags().StringVar(&role, "role", "", "Role of the user. Available roles are owner, admin, member (required).")
	_ = teamUserUpdateCmd.MarkFlagRequired("role")

	// Retrieve, delete
	flagUserId(teamUserRetrieveCmd)
	flagUserId(teamUserDeleteCmd)

}

func flagUserId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&userId, "user-id", 0, "A unique identifier of the user (required).")
	_ = cmd.MarkFlagRequired("user-id")
}
