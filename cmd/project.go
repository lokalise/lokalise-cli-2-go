package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/lokalise/go-lokalise-api/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	projectId string
	// List projects options
	includeStatistics uint8
	includeSettings   uint8
	filterTeamID      int64
	filterNames       string

	newProject     lokalise.NewProject
	newProjectLang string

	updateProject lokalise.UpdateProject
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage projects",
	Long:  "Lokalise is a project-based translation management system. We recommend to keep all platform keys in the same project. Upload iOS, Android, frontend, backend and API language files - everything that relates to a certain app or website - to the same project. We provide you with key merging or referencing options, that let you avoid recurring work for translators.",
}

var projectCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a project",
	Long:  "Creates a new project in the specified team. Requires Admin role in the team.",
	RunE: func(*cobra.Command, []string) error {
		err := json.Unmarshal([]byte(newProjectLang), &newProject.Languages)
		if err != nil {
			return err
		}

		p := Api.Projects()
		resp, err := p.Create(newProject)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var projectListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects",
	Long:  "Retrieves a list of projects available to the user, authorized with a token.",
	RunE: func(*cobra.Command, []string) error {
		p := Api.Projects()
		opts := lokalise.ProjectListOptions{
			Limit:           p.ListOpts().Limit,
			IncludeSettings: fmt.Sprintf("%d", includeSettings),
			IncludeStat:     fmt.Sprintf("%d", includeStatistics),
			FilterTeamID:    filterTeamID,
			FilterNames:     filterNames,
		}

		return repeatableList(
			func(page int64) {
				opts.Page = uint(page)
				p.SetListOptions(opts)
			},
			func() (lokalise.PageCounter, error) {
				return p.List()
			},
		)
	},
}

var projectRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieve a project",
	Long:  "Retrieves a project.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Projects().Retrieve(projectId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var projectUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a project",
	Long:  "Updates the details of a project. Requires Manage settings admin right.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Projects().Update(projectId, updateProject)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var projectDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a project",
	Long:  "Deletes a project.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Projects().Delete(projectId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var projectEmptyCmd = &cobra.Command{
	Use:   "empty",
	Short: "Empty a project",
	Long:  "Deletes all keys and translations from the project. Requires Manage settings admin right.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Projects().Truncate(projectId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	projectCmd.AddCommand(projectCreateCmd, projectListCmd, projectRetrieveCmd, projectUpdateCmd,
		projectEmptyCmd, projectDeleteCmd)
	rootCmd.AddCommand(projectCmd)

	// Create
	fs := projectCreateCmd.Flags()
	fs.StringVar(&newProject.Name, "name", "", "Name of the project (required).")
	_ = projectCreateCmd.MarkFlagRequired("name")
	fs.Int64Var(&newProject.TeamID, "team-id", 0, "ID of the team to create a project in. If this parameter is omitted, the project will be created in current team of the user, whose API token is specified.")
	fs.StringVar(&newProject.Description, "description", "", "Description of the project.")
	fs.StringVar(&newProjectLang, "languages", "", "List of languages to add (JSON, see https://lokalise.com/api2docs/curl/#transition-create-a-project-post).")
	fs.StringVar(&newProject.BaseLangISO, "base-lang-iso", "", "Language/locale code of the project base language. Should be in a scope of languages list. Use custom_iso code in case it was defined.")
	fs.StringVar(&newProject.ProjectType, "project-type", "", "Project type. Allowed values are localization_files, paged_documents.")

	// List
	fs = projectListCmd.Flags()
	fs.Uint8Var(&includeStatistics, "include-statistics", 1, "Whether to include project statistics.")
	fs.Uint8Var(&includeSettings, "include-settings", 1, "Whether to include project settings.")
	fs.Int64Var(&filterTeamID, "filter-team-id", 0, "Limit results to team ID.")
	fs.StringVar(&filterNames, "filter-names", "", "One or more project names to filter by.")

	// Retrieve
	flagProjectId(projectRetrieveCmd, false)

	// Update
	flagProjectId(projectUpdateCmd, false)
	fs = projectUpdateCmd.Flags()
	fs.StringVar(&updateProject.Name, "name", "", "Name of the project.")
	fs.StringVar(&updateProject.Description, "description", "", "Description of the project.")

	// Empty, delete
	flagProjectId(projectEmptyCmd, false)
	flagProjectId(projectDeleteCmd, false)
}

func flagProjectId(cmd *cobra.Command, isPersistent bool) {
	var fs *pflag.FlagSet
	var defaultPID string

	if isPersistent {
		fs = cmd.PersistentFlags()
		defaultPID = viper.GetString("project-id")
	} else {
		fs = cmd.Flags()
		defaultPID = ""
	}

	fs.StringVar(&projectId, "project-id", defaultPID, "Unique project identifier (required).")
	_ = cmd.MarkFlagRequired("project-id")
}
