package cmd

import (
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

var (
	teamId    int64
	userId    int64
	keyId     int64
	commentId int64
	projectId string

	role string
)

// teamUserCmd represents the team-user command
var teamUserCmd = &cobra.Command{
	Use:   "team-user",
	Short: "The Team user object",
}

// teamUserListCmd represents team-user list command
var teamUserListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all team users. Requires Admin role in the team.",
	RunE: func(cmd *cobra.Command, args []string) error {

		resp, err := Api.TeamUsers().List(teamId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

// teamUserRetrieveCmd represents team-user retrieve command
var teamUserRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a Team user object. Requires Admin role in the team.",
	RunE: func(cmd *cobra.Command, args []string) error {

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
	Short: "Updates the role of a team user. Requires Admin role in the team.",
	Long: `
Available roles are  owner, admin, member.

`,
	RunE: func(cmd *cobra.Command, args []string) error {

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
	Short: "Deletes the role of a team user. Requires Admin role in the team.",
	RunE: func(cmd *cobra.Command, args []string) error {

		resp, err := Api.TeamUsers().Delete(teamId, userId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	teamUserCmd.AddCommand(teamUserListCmd)
	teamUserCmd.AddCommand(teamUserRetrieveCmd)
	teamUserCmd.AddCommand(teamUserUpdateCmd)
	teamUserCmd.AddCommand(teamUserDeleteCmd)

	rootCmd.AddCommand(teamUserCmd)

	teamUserCmd.PersistentFlags().Int64Var(&teamId, "team-id", 0, "A unique team identifier (required)")
	_ = teamUserCmd.MarkPersistentFlagRequired("team-id")

	teamUserRetrieveCmd.Flags().Int64Var(&userId, "user-id", 0, "A unique identifier of the user (required)")
	_ = teamUserRetrieveCmd.MarkFlagRequired("user-id")

	teamUserUpdateCmd.Flags().Int64Var(&userId, "user-id", 0, "A unique identifier of the user (required)")
	_ = teamUserUpdateCmd.MarkFlagRequired("user-id")
	teamUserUpdateCmd.Flags().StringVar(&role, "role", "", "Role of the user. Available roles are owner, admin, member (required)")
	_ = teamUserUpdateCmd.MarkFlagRequired("role")

	teamUserDeleteCmd.Flags().Int64Var(&userId, "user-id", 0, "A unique identifier of the user (required)")
	_ = teamUserDeleteCmd.MarkFlagRequired("user-id")
}
