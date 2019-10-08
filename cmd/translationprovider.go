package cmd

import (
	"github.com/spf13/cobra"
)

var (
	providerId int64
)

// providerCmd represents the translation-provider command
var providerCmd = &cobra.Command{
	Use: "translation-provider",
}

var providerListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists project translation statuses",
	RunE: func(cmd *cobra.Command, args []string) error {

		resp, err := Api.TranslationProviders().List(teamId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var providerRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a translationProvider ",
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
	providerRetrieveCmd.Flags().Int64Var(&providerId, "provider-id", 0, "A unique identifier of translationProvider (required)")
	_ = providerRetrieveCmd.MarkFlagRequired("provider-id")
}
