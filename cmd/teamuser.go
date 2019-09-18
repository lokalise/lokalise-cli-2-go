package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

var (
	teamId int64
	userId int64
	role   string
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

		c, err := lokalise.NewClient(Token)
		if err != nil {
			return err
		}

		resp, err := c.TeamUsers.List(context.Background(), teamId, lokalise.PageOptions{})
		if err != nil {
			return err
		}
		output, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			return err
		}

		fmt.Println(string(output))
		return nil
	},
}

// teamUserRetrieveCmd represents team-user retrieve command
var teamUserRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a Team user object. Requires Admin role in the team.",
	RunE: func(cmd *cobra.Command, args []string) error {

		c, err := lokalise.NewClient(Token)
		if err != nil {
			return err
		}

		resp, err := c.TeamUsers.Retrieve(context.Background(), teamId, userId)
		if err != nil {
			return err
		}
		output, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			return err
		}

		fmt.Println(string(output))
		return nil
	},
}

// teamUserUpdateCmd represents team-user update command
var teamUserUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates the role of a team user. Requires Admin role in the team.",
	Long: `
Available roles are  owner, admin, member.

* owner может управлять ролями owner, admin, member, в том числе понизить привилегии
* admin может управлять ролями admin, member, НЕ может повысить привилегии до owner
* member НЕ может менять роли
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		c, err := lokalise.NewClient(Token)
		if err != nil {
			return err
		}

		resp, err := c.TeamUsers.UpdateRole(context.Background(), teamId, userId, lokalise.TeamUserRole(role))
		if err != nil {
			return err
		}
		output, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			return err
		}

		fmt.Println(string(output))
		return nil
	},
}

func init() {
	teamUserCmd.AddCommand(teamUserListCmd)
	teamUserCmd.AddCommand(teamUserRetrieveCmd)
	teamUserCmd.AddCommand(teamUserUpdateCmd)

	rootCmd.AddCommand(teamUserCmd)

	teamUserCmd.PersistentFlags().Int64Var(&teamId, "team-id", 0, "A unique team identifier (required)")
	_ = teamUserCmd.MarkPersistentFlagRequired("team-id")

	teamUserRetrieveCmd.Flags().Int64Var(&userId, "user-id", 0, "A unique identifier of the user (required)")
	_ = teamUserRetrieveCmd.MarkFlagRequired("user-id")

	teamUserUpdateCmd.Flags().Int64Var(&userId, "user-id", 0, "A unique identifier of the user (required)")
	_ = teamUserUpdateCmd.MarkFlagRequired("user-id")
	teamUserUpdateCmd.Flags().StringVar(&role, "role", "", "Role of the user. Available roles are owner, admin, member (required)")
	_ = teamUserUpdateCmd.MarkFlagRequired("role")
}
