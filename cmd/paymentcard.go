package cmd

import (
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

var (
	cardId  int64
	newCard lokalise.CreatePaymentCard
)

// cardCmd represents the payment-card command
var cardCmd = &cobra.Command{
	Use: "payment-card",
}

var cardListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists user cards",
	RunE: func(*cobra.Command, []string) error {

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
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.PaymentCards().Create(newCard)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var cardRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a card ",
	RunE: func(*cobra.Command, []string) error {

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
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.PaymentCards().Delete(cardId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	cardCmd.AddCommand(cardListCmd, cardCreateCmd, cardRetrieveCmd, cardDeleteCmd)
	rootCmd.AddCommand(cardCmd)

	// general flags

	// Create
	fs := cardCreateCmd.Flags()
	fs.StringVar(&newCard.Number, "number", "", "")
	_ = cardCreateCmd.MarkFlagRequired("number")
	fs.StringVar(&newCard.CVC, "cvc", "", "")
	_ = cardCreateCmd.MarkFlagRequired("cvc")
	fs.Int64Var(&newCard.ExpMonth, "exp-month", 0, "")
	_ = cardCreateCmd.MarkFlagRequired("exp-month")
	fs.Int64Var(&newCard.ExpYear, "exp-year", 0, "")
	_ = cardCreateCmd.MarkFlagRequired("exp-year")

	// Retrieve, delete
	flagCardId(cardRetrieveCmd)
	flagCardId(cardDeleteCmd)
}

func flagCardId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&cardId, "card-id", 0, "A unique identifier of card (required)")
	_ = cmd.MarkFlagRequired("card-id")
}
