package cmd

import (
	"encoding/json"
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

var (
	groupId        int64
	newGroup       lokalise.NewGroup
	updateGroup    lokalise.NewGroup
	groupLanguages string

	projectsList []string
	usersList    []int
)

// teamUserCmd represents the team-user command
var teamUserGroupCmd = &cobra.Command{
	Use:   "team-user-group",
	Short: "Manage team user groups",
}

// teamUserListCmd represents team-user list command
var teamUserGroupListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all user groups",
	RunE: func(*cobra.Command, []string) error {
		c := Api.TeamUserGroups()
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

var teamUserGroupCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new user group",
	Long:  "Creates a group in the team. Requires Admin right in the team.",
	RunE: func(*cobra.Command, []string) error {
		// preparing options
		if groupLanguages != "" {
			err := json.Unmarshal([]byte(groupLanguages), &newGroup.Languages)
			if err != nil {
				return err
			}
		}

		resp, err := Api.TeamUserGroups().Create(teamId, newGroup)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

// teamUserRetrieveCmd represents team-user retrieve command
var teamUserGroupRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a user group",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.TeamUserGroups().Retrieve(teamId, groupId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

// teamUserUpdateCmd represents team-user update command
var teamUserGroupUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a user group",
	Long:  "Updates the properties of a group. Requires Admin right in the team",
	RunE: func(*cobra.Command, []string) error {
		// preparing options
		if groupLanguages != "" {
			err := json.Unmarshal([]byte(groupLanguages), &updateGroup.Languages)
			if err != nil {
				return err
			}
		}

		resp, err := Api.TeamUserGroups().Update(teamId, groupId, updateGroup)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var teamUserGroupAddProjectsCmd = &cobra.Command{
	Use:   "add-projects",
	Short: "Add projects to the group",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.TeamUserGroups().AddProjects(teamId, groupId, projectsList)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var teamUserGroupRemoveProjectsCmd = &cobra.Command{
	Use:   "remove-projects",
	Short: "Add projects to the group",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.TeamUserGroups().RemoveProjects(teamId, groupId, projectsList)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var teamUserGroupAddMembersCmd = &cobra.Command{
	Use:   "add-members",
	Short: "Add members to the group",
	RunE: func(*cobra.Command, []string) error {
		// preparing options
		var ul []int64

		for _, key := range usersList {
			 ul = append(ul, int64(key))
		}

		resp, err := Api.TeamUserGroups().AddMembers(teamId, groupId, ul)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var teamUserGroupRemoveMembersCmd = &cobra.Command{
	Use:   "remove-members",
	Short: "Remove members from the group",
	RunE: func(*cobra.Command, []string) error {
		// preparing options
		var ul []int64

		for _, key := range usersList {
			ul = append(ul, int64(key))
		}

		resp, err := Api.TeamUserGroups().AddMembers(teamId, groupId, ul)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

// teamUserDeleteCmd represents team-user delete command
var teamUserGroupDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a user group",
	Long:  "Deletes a group from the team. Requires Admin right in the team",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.TeamUserGroups().Delete(teamId, groupId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	teamUserGroupCmd.AddCommand(teamUserGroupListCmd, teamUserGroupCreateCmd, teamUserGroupRetrieveCmd,
		teamUserGroupUpdateCmd, teamUserGroupDeleteCmd,
		teamUserGroupAddProjectsCmd, teamUserGroupRemoveProjectsCmd, teamUserGroupAddMembersCmd, teamUserGroupRemoveMembersCmd,
	)
	rootCmd.AddCommand(teamUserGroupCmd)

	// General
	flagTeamId(teamUserGroupCmd)

	// Create
	fs := teamUserGroupCreateCmd.Flags()
	fs.StringVar(&newGroup.Name, "name", "", "Name of the group (required).")
	_ = teamUserGroupCreateCmd.MarkFlagRequired("name")
	fs.BoolVar(&newGroup.IsReviewer, "is-reviewer", false, "Whether the group has reviewer access to the project (required).")
	_ = teamUserGroupCreateCmd.MarkFlagRequired("is-reviewer")
	fs.BoolVar(&newGroup.IsAdmin, "is-admin", false, "Whether the group has Admin access to the project (required).")
	_ = teamUserGroupCreateCmd.MarkFlagRequired("is-admin")
	fs.StringSliceVar(&newGroup.AdminRights, "admin-rights", []string{}, "List of group administrative permissions. Required if group has admin rights.")
	fs.StringVar(&groupLanguages, "languages", "", "List of languages. Required if group doesn't have admin rights. JSON, see https://lokalise.com/api2docs/curl/#transition-create-a-group-post")

	// Retrieve, delete
	flagGroupId(teamUserGroupRetrieveCmd)
	flagGroupId(teamUserGroupDeleteCmd)

	// Update
	flagGroupId(teamUserGroupUpdateCmd)
	fs = teamUserGroupUpdateCmd.Flags()
	fs.StringVar(&updateGroup.Name, "name", "", "Name of the group (required).")
	_ = teamUserGroupUpdateCmd.MarkFlagRequired("name")
	fs.BoolVar(&updateGroup.IsReviewer, "is-reviewer", false, "Whether the group has reviewer access to the project (required).")
	_ = teamUserGroupUpdateCmd.MarkFlagRequired("is-reviewer")
	fs.BoolVar(&updateGroup.IsAdmin, "is-admin", false, "Whether the group has Admin access to the project (required).")
	_ = teamUserGroupUpdateCmd.MarkFlagRequired("is-admin")
	fs.StringSliceVar(&updateGroup.AdminRights, "admin-rights", []string{}, "List of group administrative permissions. Required if group has admin rights.")
	fs.StringVar(&groupLanguages, "languages", "", "List of languages. Required if group doesn't have admin rights.")

	// Add projects
	flagGroupId(teamUserGroupAddProjectsCmd)
	teamUserGroupAddProjectsCmd.Flags().
		StringSliceVar(&projectsList, "projects", []string{}, "List of project IDs to add to group (required).")
	_ = teamUserGroupAddProjectsCmd.MarkFlagRequired("projects")

	// Remove projects
	flagGroupId(teamUserGroupRemoveProjectsCmd)
	teamUserGroupRemoveProjectsCmd.Flags().
		StringSliceVar(&projectsList, "projects", []string{}, "List of project IDs to remove from group (required).")
	_ = teamUserGroupRemoveProjectsCmd.MarkFlagRequired("projects")

	// Add members
	flagGroupId(teamUserGroupAddMembersCmd)
	teamUserGroupAddMembersCmd.Flags().
		IntSliceVar(&usersList, "users", []int{}, "List of user IDs to add to group (required).")
	_ = teamUserGroupAddMembersCmd.MarkFlagRequired("users")

	// Remove members
	flagGroupId(teamUserGroupRemoveMembersCmd)
	teamUserGroupRemoveMembersCmd.Flags().
		IntSliceVar(&usersList, "users", []int{}, "List of user IDs to remove from group (required).")
	_ = teamUserGroupRemoveMembersCmd.MarkFlagRequired("users")
}

func flagGroupId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&groupId, "group-id", 0, "A unique identifier of the group (required).")
	_ = cmd.MarkFlagRequired("group-id")
}
