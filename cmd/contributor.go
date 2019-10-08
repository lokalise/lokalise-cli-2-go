package cmd

import (
	"encoding/json"
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

var (
	contributorId        int64
	contributorCreate    lokalise.NewContributor
	permissionUpdate     lokalise.Permission
	contributorLanguages string
)

// contributorCmd represents the contributor command
var contributorCmd = &cobra.Command{
	Use:   "contributor",
	Short: "The Contributor object",
}

var contributorListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all contributors",
	RunE: func(*cobra.Command, []string) error {

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
	RunE: func(*cobra.Command, []string) error {
		// fill opts
		if contributorLanguages != "" {
			var ls []lokalise.Language
			err := json.Unmarshal([]byte(contributorLanguages), &ls)
			if err != nil {
				return err
			}
			contributorCreate.Languages = ls
		}

		resp, err := Api.Contributors().Create(projectId, []lokalise.NewContributor{contributorCreate})
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var contributorRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a contributor by its id",
	RunE: func(*cobra.Command, []string) error {

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
	RunE: func(*cobra.Command, []string) error {
		// Fill permission langs
		if contributorLanguages != "" {
			var ls []lokalise.Language
			err := json.Unmarshal([]byte(contributorLanguages), &ls)
			if err != nil {
				return err
			}
			permissionUpdate.Languages = ls
		}

		resp, err := Api.Contributors().Update(projectId, userId, permissionUpdate)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var contributorDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a user from the project. Requires Manage contributors admin right.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Contributors().Delete(projectId, userId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	contributorCmd.AddCommand(contributorListCmd, contributorCreateCmd, contributorRetrieveCmd,
		contributorUpdateCmd, contributorDeleteCmd)
	rootCmd.AddCommand(contributorCmd)

	// common for all Comment cmd`s
	flagProjectId(contributorCmd, true)

	// Create
	fs := contributorCreateCmd.Flags()
	fs.StringVar(&contributorCreate.Email, "email", "", "")
	_ = contributorCreateCmd.MarkFlagRequired("email")
	fs.StringVar(&contributorCreate.Fullname, "fullname", "", "")
	fs.BoolVar(&contributorCreate.IsAdmin, "is-admin", false, "")
	fs.BoolVar(&contributorCreate.IsReviewer, "is-reviewer", false, "")
	fs.StringVar(&contributorLanguages, "languages", "", "")
	fs.StringSliceVar(&contributorCreate.AdminRights, "admin-rights", []string{}, "")

	// Update
	flagContributorId(contributorUpdateCmd)
	fs = contributorUpdateCmd.Flags()
	fs.BoolVar(&permissionUpdate.IsAdmin, "is-admin", false, "")
	fs.BoolVar(&permissionUpdate.IsReviewer, "is-reviewer", false, "")
	fs.StringVar(&contributorLanguages, "languages", "", "")
	fs.StringSliceVar(&permissionUpdate.AdminRights, "admin-rights", []string{}, "")

	// Retrieve, delete
	flagContributorId(contributorRetrieveCmd)
	flagContributorId(contributorDeleteCmd)
}

func flagContributorId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&contributorId, "contributor-id", 0, "A unique identifier of contributor (required)")
	_ = cmd.MarkFlagRequired("contributor-id")
}
