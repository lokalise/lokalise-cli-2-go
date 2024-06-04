package cmd

import (
	"github.com/lokalise/go-lokalise-api/v4"
	"github.com/spf13/cobra"
)

var (
	orderId  string
	newOrder lokalise.CreateOrder
)

// orderCmd represents the order command
var orderCmd = &cobra.Command{
	Use:   "order",
	Short: "Manage orders",
	Long:  "Lokalise offers several human-powered translation providers, that can help translating your app or website to the most popular languages. You need to list translation providers to find out their slug, tiers and possible language pairs. You need to create a payment card in order to pay for translation orders.",
}

var orderListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all orders",
	Long:  "Lists all translation orders in the team.",
	RunE: func(*cobra.Command, []string) error {
		c := Api.Orders()
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

var orderCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an order",
	Long:  "Creates a translation order. You must have admin privileges in the project you are placing an order in.",
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
	Short: "Retrieve an order",
	Long:  "Retrieves an order.",
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
	fs := orderCreateCmd.Flags()
	fs.StringVar(&newOrder.ProjectID, "project-id", "", "Project identifier. (required).")
	_ = orderCreateCmd.MarkFlagRequired("project-id")
	fs.Int64Var(&newOrder.CardID, "card-id", 0, "Card identifier that should be used for payment. (required).")
	_ = orderCreateCmd.MarkFlagRequired("card-id")
	fs.StringVar(&newOrder.Briefing, "briefing", "", "Order briefing (required).")
	_ = orderCreateCmd.MarkFlagRequired("briefing")
	fs.StringVar(&newOrder.SourceLangISO, "source-language-iso", "", "Source language code of the order (required).")
	_ = orderCreateCmd.MarkFlagRequired("source-language-iso")
	fs.StringSliceVar(&newOrder.TargetLangISOs, "target-language-isos", []string{}, "List of target languages (required).")
	_ = orderCreateCmd.MarkFlagRequired("target-language-isos")
	fs.IntSliceVar(&newOrder.Keys, "keys", []int{}, "List of keys identifiers, included in the order (required).")
	_ = orderCreateCmd.MarkFlagRequired("keys")
	fs.StringVar(&newOrder.ProviderSlug, "provider-slug", "", "Translation provider slug (required).")
	_ = orderCreateCmd.MarkFlagRequired("provider-slug")
	fs.Int64Var(&newOrder.TranslationTierID, "translation-tier", 0, "Tier of the translation. Tiers depend on provider (order).")
	_ = orderCreateCmd.MarkFlagRequired("translation-tier")
	fs.BoolVar(&newOrder.DryRun, "dry-run", false, "Return the response without actually placing an order. Useful for price estimation. The card will not be charged.")
	fs.StringVar(&newOrder.TranslationStyle, "translation-style", "", "Only for gengo provider. Available values are formal, informal, business, friendly. Defaults to friendly.")

	// Retrieve
	flagOrderId(orderRetrieveCmd)
}

func flagOrderId(cmd *cobra.Command) {
	cmd.Flags().StringVar(&orderId, "order-id", "", "A unique identifier of order (required).")
	_ = cmd.MarkFlagRequired("order-id")
}
