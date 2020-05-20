package cmd

import (
	"github.com/spf13/cobra"
)

var (
	processId string
)

// queuedProcessCmd represents the queued-process command
var queuedProcessCmd = &cobra.Command{
	Use:   "queued-process",
	Short: "Manage queued processes",
	Long:  "Manage queued processes. Some heavy actions are processed asynchronously. Queued processes contain current processing status and eventually, final result.",
}

// queuedProcessListCmd represents queued-process list command
var queuedProcessListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all queued processes for a project",
	RunE: func(*cobra.Command, []string) error {
		resp, err := Api.QueuedProcesses().List(projectId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

// queuedProcessRetrieveCmd represents queued-process retrieve command
var queuedProcessRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieve a process for a project",
	RunE: func(*cobra.Command, []string) error {
		resp, err := Api.QueuedProcesses().Retrieve(projectId, processId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	queuedProcessCmd.AddCommand(queuedProcessListCmd, queuedProcessRetrieveCmd)
	rootCmd.AddCommand(queuedProcessCmd)

	// general flags
	flagProjectId(queuedProcessCmd, true)

	// List
	queuedProcessRetrieveCmd.Flags().StringVar(&processId, "process-id", "", "A unique identifier of the process (required).")
	_ = queuedProcessRetrieveCmd.MarkFlagRequired("process-id")
}
