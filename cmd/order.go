package cmd

import (
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

var (
	orderId  int64
	newOrder lokalise.CreateOrder
)

// orderCmd represents the order command
var orderCmd = &cobra.Command{
	Use:   "order",
	Short: "The ...",
}

var orderListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists project orders",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Orders().List(teamId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var orderCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a order in the project",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Orders().Create(teamId, newOrder)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var orderRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a order ",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Orders().Retrieve(teamId, orderId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	orderCmd.AddCommand(orderListCmd, orderCreateCmd, orderRetrieveCmd)
	rootCmd.AddCommand(orderCmd)

	// general flags
	flagTeamId(orderCmd)

	// Create
	flagCardId(orderCreateCmd)
	fs := orderCreateCmd.Flags()
	fs.StringVar(&newOrder.Briefing, "briefing", "", "")
	_ = orderCreateCmd.MarkFlagRequired("briefing")
	fs.StringVar(&newOrder.SourceLangISO, "source-language-iso", "", "")
	_ = orderCreateCmd.MarkFlagRequired("source-language-iso")
	fs.StringSliceVar(&newOrder.TargetLangISOs, "target-language-isos", []string{}, "")
	_ = orderCreateCmd.MarkFlagRequired("target-language-isos")
	fs.StringSliceVar(&newOrder.Keys, "keys", []string{}, "")
	_ = orderCreateCmd.MarkFlagRequired("keys")
	fs.StringVar(&newOrder.ProviderSlug, "provider-slug", "", "")
	_ = orderCreateCmd.MarkFlagRequired("provider-slug")
	fs.Int64Var(&newOrder.TranslationTierID, "translation-tier", 0, "")
	_ = orderCreateCmd.MarkFlagRequired("translation-tier")
	fs.BoolVar(&newOrder.DryRun, "dry-run", false, "")
	fs.StringVar(&newOrder.TranslationStyle, "translation-style", "", "")

	// Retrieve
	flagOrderId(orderRetrieveCmd)
}

func flagOrderId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&orderId, "order-id", 0, "A unique identifier of order (required)")
	_ = cmd.MarkFlagRequired("order-id")
}
