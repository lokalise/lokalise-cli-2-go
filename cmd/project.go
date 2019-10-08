package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
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
}

var projectCreateCmd = &cobra.Command{
	Use: "create",
	RunE: func(*cobra.Command, []string) error {
		err := json.Unmarshal([]byte(newProjectLang), &newProject.Languages)
		if err != nil {
			return err
		}

		p := Api.Projects()
		p.SetDebug(true)
		resp, err := p.Create(newProject)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var projectListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all projects",
	RunE: func(*cobra.Command, []string) error {

		p := Api.Projects()
		p.SetDebug(true)
		opts := lokalise.ProjectListOptions{
			IncludeSettings: fmt.Sprintf("%d", includeSettings),
			IncludeStat:     fmt.Sprintf("%d", includeStatistics),
			FilterTeamID:    filterTeamID,
			FilterNames:     filterNames,
		}

		resp, err := p.WithListOptions(opts).List()
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var projectRetrieveCmd = &cobra.Command{
	Use: "retrieve",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Projects().Retrieve(projectId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var projectUpdateCmd = &cobra.Command{
	Use: "update",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Projects().Update(projectId, updateProject)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var projectDeleteCmd = &cobra.Command{
	Use: "delete",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Projects().Delete(projectId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var projectEmptyCmd = &cobra.Command{
	Use: "empty",
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
	fs.StringVar(&newProject.Name, "name", "", "Name of the project")
	_ = projectCreateCmd.MarkFlagRequired("name")
	fs.Int64Var(&newProject.TeamID, "team-id", 0, "ID of the team to create a project in")
	fs.StringVar(&newProject.Description, "description", "", "Description of the project")
	fs.StringVar(&newProjectLang, "languages", "", "List of languages to add")
	fs.StringVar(&newProject.BaseLangISO, "base-lang-iso", "", "Language/locale code of the project base language")
	fs.StringVar(&newProject.ProjectType, "project-type", "", "Project type")

	// List
	fs = projectListCmd.Flags()
	fs.Uint8Var(&includeStatistics, "include-statistics", 1, "Whether to include project statistics")
	fs.Uint8Var(&includeSettings, "include-settings", 1, "Whether to include project settings")
	fs.Int64Var(&filterTeamID, "filter-team-id", 0, "Limit results to team ID")
	fs.StringVar(&filterNames, "filter-names", "", "One or more project names to filter by (comma separated)")

	// Retrieve
	flagProjectId(projectRetrieveCmd, false)

	// Update
	flagProjectId(projectUpdateCmd, false)
	fs = projectUpdateCmd.Flags()
	fs.StringVar(&updateProject.Name, "name", "", "Name of the project")
	fs.StringVar(&updateProject.Description, "description", "", "Description of the project")

	// Empty, delete
	flagProjectId(projectEmptyCmd, false)
	flagProjectId(projectDeleteCmd, false)
}

func flagProjectId(cmd *cobra.Command, isPersistent bool) {
	if isPersistent {
		cmd.PersistentFlags().StringVar(&projectId, "project-id", "", "A unique project identifier (required)")
		_ = cmd.MarkPersistentFlagRequired("project-id")
	} else {
		cmd.Flags().StringVar(&projectId, "project-id", "", "A unique project identifier (required)")
		_ = cmd.MarkFlagRequired("project-id")
	}
}
