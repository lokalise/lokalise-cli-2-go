package cmd

import (
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

var (
	taskId int64
)

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use: "task",
}

var taskListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists project tasks",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Tasks().List(projectId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var taskCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a task in the project",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Tasks().Create(projectId, lokalise.CreateTask{})
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var taskRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a task ",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Tasks().Retrieve(projectId, taskId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var taskUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates the properties of a task",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Tasks().Update(projectId, taskId, lokalise.UpdateTask{})
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var taskDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a task from the project.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Tasks().Delete(projectId, taskId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	taskCmd.AddCommand(taskListCmd, taskCreateCmd, taskRetrieveCmd, taskUpdateCmd, taskDeleteCmd)
	rootCmd.AddCommand(taskCmd)

	// general flags
	flagProjectId(taskCmd, true)

	// separate flags for every command
	flagTaskId(taskRetrieveCmd)
	flagTaskId(taskUpdateCmd)
	flagTaskId(taskDeleteCmd)
}

func flagTaskId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&taskId, "task-id", 0, "A unique identifier of task (required)")
	_ = cmd.MarkFlagRequired("task-id")
}
