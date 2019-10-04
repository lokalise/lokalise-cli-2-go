package cmd

import (
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

var (
	orderId int64
)

// orderCmd represents the order command
var orderCmd = &cobra.Command{
	Use:   "order",
	Short: "The ...",
}

var orderListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists project orders",
	RunE: func(cmd *cobra.Command, args []string) error {
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

	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Orders().Create(teamId, lokalise.CreateOrder{}) // fixme

		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var orderRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a order ",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Orders().Retrieve(teamId, orderId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	orderCmd.AddCommand(orderListCmd)
	orderCmd.AddCommand(orderCreateCmd)
	orderCmd.AddCommand(orderRetrieveCmd)

	rootCmd.AddCommand(orderCmd)

	// general flags
	withProjectId(orderCmd, true)

	// separate flags for every command
	flagOrderId(orderRetrieveCmd)
}

func flagOrderId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&orderId, "order-id", 0, "A unique identifier of order (required)")
	_ = cmd.MarkFlagRequired("order-id")
}
