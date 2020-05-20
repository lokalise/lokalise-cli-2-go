package cmd

import (
	"github.com/lokalise/go-lokalise-api/v3"
	"github.com/spf13/cobra"
)

var (
	cardId  int64
	newCard lokalise.CreatePaymentCard
)

// cardCmd represents the payment-card command
var cardCmd = &cobra.Command{
	Use:   "payment-card",
	Short: "Manage payment cards",
	Long:  "Credit cards are used to pay for translation orders. Each user has their own cards, that are not shared with other users. We do not store credit card details. Once the card is added, we send the details to Stripe and receive the card token, which can only be used for order purchases at Lokalise.",
}

var cardListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all cards",
	Long:  "Lists all user payment cards.",
	RunE: func(*cobra.Command, []string) error {
		c := Api.PaymentCards()
		pageOpts := c.PageOpts()

		return repeatableList(
			func(p int64) {
				pageOpts.Page = uint(p)
				c.SetPageOptions(pageOpts)
			},
			func() (lokalise.PageCounter, error) {
				return c.List()
			},
		)
	},
}

var cardCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a card",
	Long:  "Adds new payment card to user cards.",
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
	Short: "Retrieve a card",
	Long:  "Retrieves a payment card.",
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
	Short: "Delete a card",
	Long:  "Deletes a payment card from user cards.",
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
	fs.StringVar(&newCard.Number, "number", "", "Card number (required).")
	_ = cardCreateCmd.MarkFlagRequired("number")
	fs.StringVar(&newCard.CVC, "cvc", "", "3-digit card CVC code (required).")
	_ = cardCreateCmd.MarkFlagRequired("cvc")
	fs.Int64Var(&newCard.ExpMonth, "exp-month", 0, "Card expiration month (1-12) (required).")
	_ = cardCreateCmd.MarkFlagRequired("exp-month")
	fs.Int64Var(&newCard.ExpYear, "exp-year", 0, "Card expiration year (required).")
	_ = cardCreateCmd.MarkFlagRequired("exp-year")

	// Retrieve, delete
	flagCardId(cardRetrieveCmd)
	flagCardId(cardDeleteCmd)
}

func flagCardId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&cardId, "card-id", 0, "A unique identifier of the card (required).")
	_ = cmd.MarkFlagRequired("card-id")
}
