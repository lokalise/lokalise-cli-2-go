package cmd

import (
	"github.com/spf13/cobra"
)

var (
	providerId int64
)

// providerCmd represents the translation-provider command
var providerCmd = &cobra.Command{
	Use:   "translation-provider",
	Short: "The ...",
}

var providerListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists project translation statuses",
	/*RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.TranslationProviders().List(teamId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},*/
}

var providerRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a translationProvider ",
	/*RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.TranslationProviders().Retrieve(teamId, providerId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},*/
}

func init() {
	providerCmd.AddCommand(providerListCmd)
	providerCmd.AddCommand(providerRetrieveCmd)

	rootCmd.AddCommand(providerCmd)

	// general flags
	// withTeamId(providerCmd, true) todo implement

	// separate flags for every command
	providerCmd.Flags().Int64Var(&providerId, "provider-id", 0, "A unique identifier of translationProvider (required)")
	_ = providerCmd.MarkFlagRequired("provider-id")
}
