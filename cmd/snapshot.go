package cmd

import (
	"github.com/spf13/cobra"
)

var (
	snapshotId    int64
	snapshotTitle string
)

// snapshotCmd represents the snapshot command
var snapshotCmd = &cobra.Command{
	Use: "snapshot",
}

var snapshotListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists project snapshots",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Snapshots().List(projectId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var snapshotCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates snapshot of the project",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Snapshots().Create(projectId, snapshotTitle)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var snapshotRestoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restores project snapshot to a project copy",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Snapshots().Restore(projectId, snapshotId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var snapshotDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a snapshot from the project.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Snapshots().Delete(projectId, snapshotId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	snapshotCmd.AddCommand(snapshotListCmd, snapshotCreateCmd, snapshotRestoreCmd, snapshotDeleteCmd)
	rootCmd.AddCommand(snapshotCmd)

	// general flags
	flagProjectId(snapshotCmd, true)

	// separate flags for every command
	snapshotCreateCmd.Flags().StringVar(&snapshotTitle, "title", "", "Set snapshot title")

	flagSnapshotId(snapshotDeleteCmd)
	flagSnapshotId(snapshotRestoreCmd)
}

func flagSnapshotId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&snapshotId, "snapshot-id", 0, "A unique identifier of snapshot (required)")
	_ = cmd.MarkFlagRequired("snapshot-id")
}
