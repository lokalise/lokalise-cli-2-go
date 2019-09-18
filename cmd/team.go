package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

// teamCmd represents the team command
var teamCmd = &cobra.Command{
	Use:   "team",
	Short: "The Team object",
}

// teamListCmd represents team list command
var teamListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all teams available to the user",
	RunE: func(cmd *cobra.Command, args []string) error {

		c, err := lokalise.NewClient(Token)
		if err != nil {
			return err
		}

		resp, err := c.Teams.List(context.Background(), lokalise.PageOptions{})
		if err != nil {
			return err
		}

		output, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			return err
		}

		fmt.Println(string(output))
		return nil
	},
}

func init() {
	teamCmd.AddCommand(teamListCmd)

	rootCmd.AddCommand(teamCmd)
}
