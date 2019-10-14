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
	Use:   "comment",
	Short: "Manage key comments",
	Long:  "Comments can be used to give translators a context about the key, or as a discussion about certain aspects of translation for the key. There is a separate comments thread for each key. All comments are cross-posted into project chat.",
}

var commentListCmd = &cobra.Command{
	Use:   "list",
	Short: "List project comments",
	Long:  "Retrieves a list of all comments in the project.",
	RunE: func(*cobra.Command, []string) error {
		c := Api.Comments()
		pageOpts := c.PageOpts()

		return repeatableList(
			func(p int64) {
				pageOpts.Page = uint(p)
				c.SetPageOptions(pageOpts)
			},
			func() (lokalise.PageCounter, error) {
				return c.ListProject(projectId)
			},
		)
	},
}

var commentListKeyCmd = &cobra.Command{
	Use:   "list-key",
	Short: "List key comments",
	Long:  "Retrieves a list of all comments of the key.",
	RunE: func(*cobra.Command, []string) error {
		c := Api.Comments()
		pageOpts := c.PageOpts()

		return repeatableList(
			func(p int64) {
				pageOpts.Page = uint(p)
				c.SetPageOptions(pageOpts)
			},
			func() (lokalise.PageCounter, error) {
				return c.ListByKey(projectId, keyId)
			},
		)
	},
}

var commentCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Add a comment",
	Long:  "Adds a comment to the skey.",
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
	Short: "Retrieve a comment",
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
	Short: "Delete a comment",
	Long:  "Deletes a comment from the project. Authenticated user can only delete own comments.",
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
	commentCreateCmd.Flags().StringVar(&comment, "comment", "", "The comment to add (required).")
	_ = commentCreateCmd.MarkFlagRequired("comment")

	// Retrieve
	flagKeyId(commentRetrieveCmd)
	flagCommentId(commentRetrieveCmd)

	// Delete
	flagKeyId(commentDeleteCmd)
	flagCommentId(commentDeleteCmd)
}

func flagCommentId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&commentId, "comment-id", 0, "A unique identifier of comment (required).")
	_ = cmd.MarkFlagRequired("comment-id")
}
