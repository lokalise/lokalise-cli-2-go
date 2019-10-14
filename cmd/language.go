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
	Use:   "language",
	Short: "Manage languages",
	Long: `There are over 500 predefined language/dialect combinations available in Lokalise. In case you require a custom language/dialect combination use custom_X languages (where X is a number from 1 to 100). You may override language code and name when adding a language, or update an existing language properties later.

There are situations when it is necessary to export different language codes to different platforms (e.g. zh-Hans to iOS and zh_Hans to Web). In such cases you need to set any preferred version and ise export parameter to set language mapping depending on the file format.
`,
}

var languageListCmd = &cobra.Command{
	Use:   "list",
	Short: "List project languages",
	Long:  "Retrieves a list of project languages.",
	RunE: func(*cobra.Command, []string) error {
		c := Api.Languages()
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

var languageListSystemCmd = &cobra.Command{
	Use:   "list-system",
	Short: "List system languages",
	Long:  "Retrieves a list of system languages.s",
	RunE: func(*cobra.Command, []string) error {
		c := Api.Languages()
		pageOpts := c.PageOpts()

		return repeatableList(
			func(p int64) {
				pageOpts.Page = uint(p)
				c.SetPageOptions(pageOpts)
			},
			func() (lokalise.PageCounter, error) {
				return c.ListSystem()
			},
		)
	},
}

var languageCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create languages",
	Long: `Creates one or more languages in the project. Requires Manage languages admin right.

	The language_iso is the identifer of one of the system languages. You are only required to include the language_iso attribute, however you may override the default language code, language name and plural forms as well.
`,
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
	Short: "Retrieve a language",
	Long:  "Retrieves a language.",
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
	Short: "Update a language",
	Long:  "Updates the properties of a language. Requires Manage languages admin right.",
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
	Short: "Delete a language.",
	Long:  "Deletes a language from the project. Requires Manage languages admin right.",
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
	fs.StringVar(&newLanguage.LangISO, "lang-iso", "", "A unique language code in the system.")
	_ = languageCreateCmd.MarkFlagRequired("lang-iso")
	fs.StringVar(&newLanguage.CustomISO, "custom-iso", "", "Override language/locale code.")
	fs.StringVar(&newLanguage.CustomName, "custom-name", "", "Override language name.")
	fs.StringSliceVar(&newLanguage.CustomPluralForms, "custom-plural-forms", []string{}, "Override list of supported plural forms for this language.")

	// Update
	flagLangId(languageUpdateCmd)
	fs = languageUpdateCmd.Flags()
	fs.StringVar(&updateLanguage.LangISO, "lang-iso", "", "Language/locale code.")
	fs.StringVar(&updateLanguage.LangName, "lang-name", "", "Language name.")
	fs.StringSliceVar(&updateLanguage.PluralForms, "plural-forms", []string{}, "List of supported plural forms.")

	// Retrieve & delete
	flagLangId(languageRetrieveCmd)
	flagLangId(languageDeleteCmd)
}

func flagLangId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&languageId, "lang-id", 0, "A unique identifier of the language (required).")
	_ = cmd.MarkFlagRequired("lang-id")
}
