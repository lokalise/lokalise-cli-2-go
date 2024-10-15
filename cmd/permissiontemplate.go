package cmd

import (
	"github.com/spf13/cobra"
)

var permissionTemplateCmd = &cobra.Command{
	Use:   "permission-template",
	Short: "Manage permission templates for a team",
	Long:  "Manage all templates for permissions inside a team.",
}

var permissionTemplateListCmd = &cobra.Command{
	Use:   "list",
	Short: "List permission templates",
	Long:  "Retrieves a list of permission templates.",
	RunE: func(*cobra.Command, []string) error {
		c := Api.PermissionTemplates()
		data, err := c.ListPermissionRoles(teamId)
		if err != nil {
			return err
		}
		return printJson(data)
	},
}

func init() {
	permissionTemplateCmd.AddCommand(permissionTemplateListCmd)
	rootCmd.AddCommand(permissionTemplateCmd)

	flagTeamId(permissionTemplateCmd)
}
