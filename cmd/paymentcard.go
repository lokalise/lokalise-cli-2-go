package cmd

import (
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

var (
	cardId int64
)

// cardCmd represents the payment-card command
var cardCmd = &cobra.Command{
	Use:   "payment-card",
	Short: "The ...",
}

var cardListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists user cards",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.PaymentCards().List()
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var cardCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a card in the project",

	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.PaymentCards().Create(lokalise.CreatePaymentCard{}) // fixme

		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var cardRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a card ",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.PaymentCards().Retrieve(cardId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var cardDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a card from the project.",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.PaymentCards().Delete(cardId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	cardCmd.AddCommand(cardListCmd)
	cardCmd.AddCommand(cardCreateCmd)
	cardCmd.AddCommand(cardRetrieveCmd)
	cardCmd.AddCommand(cardDeleteCmd)

	rootCmd.AddCommand(cardCmd)

	// general flags

	// separate flags for every command
	flagCardId(cardRetrieveCmd)
	flagCardId(cardDeleteCmd)
}

func flagCardId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&cardId, "card-id", 0, "A unique identifier of card (required)")
	_ = cmd.MarkFlagRequired("card-id")
}
