package cmd

import (
	"github.com/lokalise/go-lokalise-api/v4"
	"github.com/spf13/cobra"
)

var (
	providerId int64
)

// providerCmd represents the translation-provider command
var providerCmd = &cobra.Command{
	Use:   "translation-provider",
	Short: "List translation providers",
	Long:  "Translation providers are used for translation orders.",
}

var providerListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all providers",
	Long:  "Lists all translation providers.",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := Api.TranslationProviders()
		pageOpts := c.PageOpts()

		return repeatableList(
			func(p int64) {
				pageOpts.Page = uint(p)
				c.SetPageOptions(pageOpts)
			},
			func() (lokalise.PageCounter, error) {
				return c.List(teamId)
			},
		)
	},
}

var providerRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieve a provider",
	Long:  "Retrieves a translation provider with tiers and available language pairs.",
	RunE: func(cmd *cobra.Command, args []string) error {

		resp, err := Api.TranslationProviders().Retrieve(teamId, providerId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	providerCmd.AddCommand(providerListCmd, providerRetrieveCmd)
	rootCmd.AddCommand(providerCmd)

	// general flags
	flagTeamId(providerCmd)

	// Retrieve
	providerRetrieveCmd.Flags().Int64Var(&providerId, "provider-id", 0, "A unique identifier of the translation provider (required).")
	_ = providerRetrieveCmd.MarkFlagRequired("provider-id")
}
