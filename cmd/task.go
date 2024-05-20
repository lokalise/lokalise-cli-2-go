package cmd

import (
	"encoding/json"
	"github.com/lokalise/go-lokalise-api/v4"
	"github.com/spf13/cobra"
)

var (
	taskId int64

	updateTask                lokalise.UpdateTask
	newTask                   lokalise.CreateTask
	newTaskKeys               []int
	autoCloseLang             bool
	autoCloseTask             bool
	customTranslationStatuses []int

	taskLanguages string
	taskType      string

	filterTitle string
)

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage tasks",
}

var taskListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  "Lists all tasks in the project.",
	RunE: func(*cobra.Command, []string) error {
		t := Api.Tasks()
		pageOpts := lokalise.TaskListOptions{
			FilterTitle: filterTitle,
			Limit:       t.ListOpts().Limit,
		}

		return repeatableList(
			func(p int64) {
				pageOpts.Page = uint(p)
				t.SetListOptions(pageOpts)
			},
			func() (lokalise.PageCounter, error) {
				return t.List(projectId)
			},
		)
	},
}

var taskCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a task",
	Long:  "Creates a task in the project. Requires Manage tasks admin right.",
	RunE: func(*cobra.Command, []string) error {
		// preparing options
		err := newTaskFillFields()
		if err != nil {
			return err
		}

		resp, err := Api.Tasks().Create(projectId, newTask)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var taskRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieve a task",
	Long:  "Retrieves a task.",
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
	Short: "Update a task",
	Long:  "Updates the properties of a task. Requires Manage tasks admin right.",
	RunE: func(*cobra.Command, []string) error {
		err := updateTaskFillFields()
		if err != nil {
			return err
		}

		resp, err := Api.Tasks().Update(projectId, taskId, updateTask)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var taskDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long:  "Deletes a task from the project. Requires Manage tasks admin right.",
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

	// List
	taskListCmd.Flags().StringVar(&filterTitle, "filter-title", "", "Set title filter for the list.")

	// Create
	fs := taskCreateCmd.Flags()
	fs.StringVar(&newTask.Title, "title", "", "Task title (required).")
	_ = taskCreateCmd.MarkFlagRequired("title")
	fs.StringVar(&newTask.Description, "description", "", "Short description of the task.")
	fs.StringVar(&newTask.DueDate, "due-date", "", "Due date in `Y-m-d H:i:s` format. Example: `2018-12-24 23:59:59`.")
	fs.IntSliceVar(&newTaskKeys, "keys", []int{}, "List of keys identifiers, included in task. Required if parent_task_id is not specified.")
	fs.StringVar(&taskLanguages, "languages", "", "List of languages in the task. One of `users` or `groups` must be provided (JSON, required, see https://lokalise.com/api2docs/curl/#transition-create-a-task-post).")
	_ = taskCreateCmd.MarkFlagRequired("languages")
	fs.BoolVar(&autoCloseLang, "auto-close-languages", true, "Whether languages should be closed automatically upon completion of the last item (default true). Use --auto-close-languages=false to disable.")
	fs.BoolVar(&autoCloseTask, "auto-close-task", true, "Whether the task should be automatically closed upon all language completion (default true). Use --auto-close-task=false to disable.")
	fs.StringVar(&taskType, "task-type", "", "Specify if task type is `translation` (default) or `review`.")
	fs.Int64Var(&newTask.ParentTaskID, "parent-task-id", 0, "If task_type is review, it can have a parent task. Current task will be opened when parent task is closed.")
	fs.StringSliceVar(&newTask.ClosingTags, "closing-tags", []string{}, "Tags that will be added to affected keys when task is closed.")
	fs.BoolVar(&newTask.LockTranslations, "do-lock-translations", false, "If set to 1, will lock translations for non-assigned project members.")
	fs.IntSliceVar(&customTranslationStatuses, "custom-translation-status-ids", []int{}, "IDs of custom translation statuses that will be applied to task items after item is completed.")

	// Retrieve
	flagTaskId(taskRetrieveCmd)

	// Update
	flagTaskId(taskUpdateCmd)
	fs = taskUpdateCmd.Flags()
	fs.StringVar(&updateTask.Title, "title", "", "Task title.")
	fs.StringVar(&updateTask.Description, "description", "", "Short description of the task.")
	fs.StringVar(&updateTask.DueDate, "due-date", "", "Due date in `Y-m-d H:i:s` format. Example: `2018-12-24 23:59:59`.")
	fs.StringVar(&taskLanguages, "languages", "", "List of languages to update. JSON, differs from a creation list, see https://lokalise.com/api2docs/curl/#transition-update-a-task-put.")
	fs.BoolVar(&autoCloseLang, "auto-close-languages", true, "Whether languages should be closed automatically upon completion of the last item (default true). Use --auto-close-languages=false to disable.")
	fs.BoolVar(&autoCloseTask, "auto-close-task", true, "Whether the task should be automatically closed upon all language completion (default true). Use --auto-close-task=false to disable.")
	fs.BoolVar(&updateTask.CloseTask, "close-task", false, "Whether the task should be closed and notifications sent. The task cannot be reopened again.")
	fs.StringSliceVar(&updateTask.ClosingTags, "closing-tags", []string{}, "Tags that will be added to affected keys when task is closed.")
	fs.BoolVar(&updateTask.LockTranslations, "do-lock-translations", false, "If set to 1, will lock translations for non-assigned project members.")

	// Delete
	flagTaskId(taskDeleteCmd)
}

func flagTaskId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&taskId, "task-id", 0, "A unique identifier of the task (required).")
	_ = cmd.MarkFlagRequired("task-id")
}

func newTaskFillFields() error {
	for _, key := range newTaskKeys {
		newTask.Keys = append(newTask.Keys, int64(key))
	}
	for _, statusID := range customTranslationStatuses {
		newTask.CustomTranslationStatusIDs = append(newTask.CustomTranslationStatusIDs, int64(statusID))
	}

	if taskLanguages != "" {
		err := json.Unmarshal([]byte(taskLanguages), &newTask.Languages)
		if err != nil {
			return err
		}
	}

	newTask.AutoCloseLanguages = &autoCloseLang
	newTask.AutoCloseTask = &autoCloseTask
	newTask.TaskType = lokalise.TaskType(taskType)
	return nil
}

func updateTaskFillFields() error {
	if taskLanguages != "" {
		err := json.Unmarshal([]byte(taskLanguages), &updateTask.Languages)
		if err != nil {
			return err
		}
	}

	updateTask.AutoCloseLanguages = &autoCloseLang
	updateTask.AutoCloseTask = &autoCloseTask
	return nil
}
