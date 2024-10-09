package cmd

import (
	"encoding/json"

	"github.com/lokalise/go-lokalise-api/v4"
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
	Short: "Manage project contributors",
	Long:  "You may add unlimited number of contributors to your project. User roles include admin, translator and reviewer.",
}

var contributorListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all contributors",
	Long:  "Lists contributors of the project, including access levels to the project languages. Admins always have read/write access to all languages.",
	RunE: func(*cobra.Command, []string) error {
		c := Api.Contributors()
		pageOpts := c.PageOpts()

		return repeatableList(
			func(p int64) {
				pageOpts.Page = uint(p)
				c.SetPageOptions(pageOpts)
			},
			func() (lokalise.PageCounter, error) {
				return c.List(projectId)
			},
		)
	},
}

var contributorCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a contributor",
	Long: `Creates a contributor in the project.
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
	Short: "Retrieve a contributor",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Contributors().Retrieve(projectId, contributorId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var contributorUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a contributor",
	Long: `Updates a contributor.
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

		resp, err := Api.Contributors().Update(projectId, contributorId, permissionUpdate)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var contributorDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a contributor",
	Long:  "Deletes a user from the project. Requires Manage contributors admin right.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Contributors().Delete(projectId, contributorId)
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
	fs.StringVar(&contributorCreate.Email, "email", "", "E-mail (required).")
	_ = contributorCreateCmd.MarkFlagRequired("email")
	fs.StringVar(&contributorCreate.Fullname, "fullname", "", "Full name (only valid for inviting users, who previously did not have an account in Lokalise).")
	fs.BoolVar(&contributorCreate.IsAdmin, "is-admin", false, "Whether the user has Admin access to the project.")
	_ = fs.MarkDeprecated("is-admin", "--is-admin is deprecated and will be removed. Use admin-rights to set permissions to the user.")
	fs.BoolVar(&contributorCreate.IsReviewer, "is-reviewer", false, "Whether the user has Reviewer access to the project.")
	_ = fs.MarkDeprecated("is-reviewer", "--is-reviewer is deprecated and will be removed. Use admin-rights to set permissions to the user.")
	fs.Int64Var(&contributorCreate.RoleId, "role-id", 0, "Permission template id for the contributor. By setting this admin_rights will be ignored and a template will be assigned with predefined permission set.")
	fs.StringVar(&contributorLanguages, "languages", "", "List of languages, accessible to the user. Required if is_admin is set to false (JSON, see https://lokalise.com/api2docs/curl/#transition-create-contributors-post).")
	fs.StringSliceVar(&contributorCreate.AdminRights, "admin-rights", []string{}, "Custom list of user permissions. Possible values are activity, contributors, branches_create, branches_main_modify, branches_merge, custom_status_modify, download, glossary, glossary_edit, glossary_delete, keys, manage_languages, review, screenshots, settings, statistics, tasks, upload. Omitted or empty parameter will set no rights for the user.")

	// Update
	flagContributorId(contributorUpdateCmd)
	fs = contributorUpdateCmd.Flags()
	fs.BoolVar(&permissionUpdate.IsAdmin, "is-admin", false, "Whether the user has Admin access to the project. Deprecated.")
	_ = fs.MarkDeprecated("is-admin", "--is-admin is deprecated and will be removed. Use admin-rights to set permissions to the user.")
	fs.BoolVar(&permissionUpdate.IsReviewer, "is-reviewer", false, "Whether the user has Reviewer access to the project. Deprecated.")
	_ = fs.MarkDeprecated("is-reviewer", "--is-reviewer is deprecated and will be removed. Use admin-rights to set permissions to the user.")
	fs.StringVar(&contributorLanguages, "languages", "", "List of languages, accessible to the user (JSON, see https://lokalise.com/api2docs/curl/#transition-update-a-contributor-put).")
	fs.StringSliceVar(&permissionUpdate.AdminRights, "admin-rights", []string{}, "Custom list of user permissions. Possible values are activity, contributors, branches_create, branches_main_modify, branches_merge, custom_status_modify, download, glossary, glossary_edit, glossary_delete, keys, manage_languages, review, screenshots, settings, statistics, tasks, upload. Empty parameter will set no rights for the user.")
	fs.Int64Var(&permissionUpdate.RoleId, "role-id", 0, "Permission template id for the contributor. By setting this admin_rights will be ignored and a template will be assigned with predefined permission set.")

	// Retrieve, delete
	flagContributorId(contributorRetrieveCmd)
	flagContributorId(contributorDeleteCmd)
}

func flagContributorId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&contributorId, "contributor-id", 0, "A unique identifier of contributor (required).")
	_ = cmd.MarkFlagRequired("contributor-id")
}
