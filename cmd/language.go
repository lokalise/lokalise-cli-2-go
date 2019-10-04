package cmd

import (
	"github.com/spf13/cobra"
)

var (
	languageId int64
)

// languageCmd represents the language command
var languageCmd = &cobra.Command{
	Use:   "language",
	Short: "The ...",
}

var languageListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists project languages",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Languages().ListProject(projectId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var languageListSystemCmd = &cobra.Command{
	Use:   "list-system",
	Short: "Lists system languages",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Languages().ListSystem()
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var languageCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a language in the project",

	/*RunE: func(cmd *cobra.Command, args []string) error {
		c := lokalise.CustomLanguage{languageOptions}
		resp, err := Api.Languages().Create(projectId, []lokalise.CustomLanguage{c})
		if err != nil {
			return err
		}
		return printJson(resp)
	},*/
}

var languageRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a language ",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Languages().Retrieve(projectId, languageId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var languageUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates the properties of a language",
	/*RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Languages().Update(projectId, languageId, language)
		if err != nil {
			return err
		}
		return printJson(resp)
	},*/
}

var languageDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a language from the project. Requires Manage languages admin right.",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Languages().Delete(projectId, languageId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	languageCmd.AddCommand(languageListCmd)
	languageCmd.AddCommand(languageListSystemCmd)
	languageCmd.AddCommand(languageCreateCmd)
	languageCmd.AddCommand(languageRetrieveCmd)
	languageCmd.AddCommand(languageUpdateCmd)
	languageCmd.AddCommand(languageDeleteCmd)

	rootCmd.AddCommand(languageCmd)

	// common for all Comment cmd`s
	withProjectId(languageCmd, true) // fixme except for listSystem

	// separate flags for every command
	flagLangId(languageCreateCmd)
	flagLangId(languageRetrieveCmd)
	flagLangId(languageDeleteCmd)
}

func flagLangId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&languageId, "lang-id", 0, "A unique identifier of language (required)")
	_ = cmd.MarkFlagRequired("lang-id")
}
