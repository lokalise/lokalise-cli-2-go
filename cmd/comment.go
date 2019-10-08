package cmd

import (
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

var (
	commentId int64
	comment   string
)

// commentCmd represents the comment command
var commentCmd = &cobra.Command{
	Use: "comment",
}

var commentListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all comments",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Comments().ListProject(projectId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var commentListKeyCmd = &cobra.Command{
	Use:   "list-key",
	Short: "Retrieves a list of all comments for a key",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Comments().ListByKey(projectId, keyId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var commentCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Adds a comment to the key",
	RunE: func(*cobra.Command, []string) error {

		c := lokalise.NewComment{Comment: comment}
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
	RunE: func(*cobra.Command, []string) error {

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
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Comments().Delete(projectId, keyId, commentId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	commentCmd.AddCommand(commentListCmd, commentListKeyCmd, commentCreateCmd, commentRetrieveCmd, commentDeleteCmd)
	rootCmd.AddCommand(commentCmd)

	// common for all Comment cmd`s
	flagProjectId(commentCmd, true)

	// List key
	flagKeyId(commentListKeyCmd)

	// Create
	flagKeyId(commentCreateCmd)
	commentCreateCmd.Flags().StringVar(&comment, "comment", "", "The comment")
	_ = commentCreateCmd.MarkFlagRequired("comment")

	// Retrieve
	flagKeyId(commentRetrieveCmd)
	flagCommentId(commentRetrieveCmd)

	// Delete
	flagKeyId(commentDeleteCmd)
	flagCommentId(commentDeleteCmd)
}

func flagCommentId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&commentId, "comment-id", 0, "A unique identifier of comment (required)")
	_ = cmd.MarkFlagRequired("comment-id")
}
