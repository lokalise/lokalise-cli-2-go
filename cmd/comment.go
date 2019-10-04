package cmd

import (
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

// commentCmd represents the comment command
var commentCmd = &cobra.Command{
	Use:   "comment",
	Short: "A brief description of your command",
}

var commentListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all comments",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Comments().ListProject(projectId, lokalise.PageOptions{}) // todo remove
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var commentListKeyCmd = &cobra.Command{
	Use:   "list-key",
	Short: "Retrieves a list of all comments for a key",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Comments().ListByKey(projectId, keyId, lokalise.PageOptions{}) // todo remove
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var commentCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Adds a comment to the key",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := lokalise.NewComment{comment}
		resp, err := Api.Comments().Create(projectId, keyId, []lokalise.NewComment{c})
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var commentRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a comment by its id",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Comments().Retrieve(projectId, keyId, commentId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var commentDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a comment from the project. Authenticated user can only delete own comments.",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Comments().Delete(projectId, keyId, commentId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	commentCmd.AddCommand(commentListCmd)
	commentCmd.AddCommand(commentListKeyCmd)
	commentCmd.AddCommand(commentCreateCmd)
	commentCmd.AddCommand(commentRetrieveCmd)
	commentCmd.AddCommand(commentDeleteCmd)

	rootCmd.AddCommand(commentCmd)

	// common for all Comment cmd`s
	commentCmd.PersistentFlags().StringVar(&projectId, "project-id", "", "A unique project identifier (required)")
	_ = commentCmd.MarkPersistentFlagRequired("project-id")

	// separate flags for every command
	withKeyId(commentCreateCmd)
	withKeyId(commentRetrieveCmd)
	withKeyId(commentDeleteCmd)
}
