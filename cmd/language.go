package cmd

import (
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

var (
	languageId     int64
	newLanguage    lokalise.NewLanguage
	updateLanguage lokalise.UpdateLanguage
)

// languageCmd represents the language command
var languageCmd = &cobra.Command{
	Use: "language",
}

var languageListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists project languages",
	RunE: func(*cobra.Command, []string) error {

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
	RunE: func(*cobra.Command, []string) error {

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
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Languages().Create(projectId, []lokalise.NewLanguage{newLanguage})
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var languageRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a language ",
	RunE: func(*cobra.Command, []string) error {

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
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Languages().Update(projectId, languageId, updateLanguage)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var languageDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a language from the project. Requires Manage languages admin right.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Languages().Delete(projectId, languageId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	languageCmd.AddCommand(languageListCmd, languageListSystemCmd, languageCreateCmd, languageRetrieveCmd,
		languageUpdateCmd, languageDeleteCmd)
	rootCmd.AddCommand(languageCmd)

	// common for all Comment cmd`s
	flagProjectId(languageCmd, true)

	// Create
	fs := languageCreateCmd.Flags()
	fs.StringVar(&newLanguage.LangISO, "lang-iso", "", "")
	_ = languageCreateCmd.MarkFlagRequired("lang-iso")
	fs.StringVar(&newLanguage.CustomISO, "custom-iso", "", "")
	fs.StringVar(&newLanguage.CustomName, "custom-name", "", "")
	fs.StringSliceVar(&newLanguage.CustomPluralForms, "custom-plural-forms", []string{}, "")

	// Update
	flagLangId(languageUpdateCmd)
	fs = languageUpdateCmd.Flags()
	fs.StringVar(&updateLanguage.LangISO, "lang-iso", "", "")
	fs.StringVar(&updateLanguage.LangName, "lang-name", "", "")
	fs.StringSliceVar(&updateLanguage.PluralForms, "plural-forms", []string{}, "")

	// Retrieve & delete
	flagLangId(languageRetrieveCmd)
	flagLangId(languageDeleteCmd)
}

func flagLangId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&languageId, "lang-id", 0, "A unique identifier of language (required)")
	_ = cmd.MarkFlagRequired("lang-id")
}
