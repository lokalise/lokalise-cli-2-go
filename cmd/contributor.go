package cmd

import (
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

// contributorCmd represents the contributor command
var contributorCmd = &cobra.Command{
	Use:   "contributor",
	Short: "The Contributor object",
}

var contributorListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all contributors",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Contributors().List(projectId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var contributorCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a contributor in the project",
	Long: `
Requires Manage contributors admin right.

If is_admin flag is set to true, the user would automatically get access to all project languages, 
overriding supplied languages object. Attribute fullname will be ignored, 
if the user has already been registered in Lokalise.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := lokalise.CustomContributor{contributorOptions}
		resp, err := Api.Contributors().Create(projectId, []lokalise.CustomContributor{c})
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var contributorRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a contributor by its id",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Contributors().Retrieve(projectId, userId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var contributorUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates the properties of a contributor",
	Long: `
Requires Manage contributors admin right.

If you want to give an existing contributor access to a new language, you must specify full languages array, 
including the previously added languages as well.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Contributors().Update(projectId, userId, contributorOptions) // fixme in lib incorrect signature
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var contributorDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a user from the project. Requires Manage contributors admin right.",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Contributors().Delete(projectId, userId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	contributorCmd.AddCommand(contributorListCmd)
	contributorCmd.AddCommand(contributorCreateCmd)
	contributorCmd.AddCommand(contributorRetrieveCmd)
	contributorCmd.AddCommand(contributorUpdateCmd)
	contributorCmd.AddCommand(contributorDeleteCmd)

	rootCmd.AddCommand(contributorCmd)

	// common for all Comment cmd`s
	withProjectId(contributorCmd, true)

	// separate flags for every command
	withKeyId(contributorCreateCmd)
	withKeyId(contributorRetrieveCmd)
	withKeyId(contributorDeleteCmd)
}
