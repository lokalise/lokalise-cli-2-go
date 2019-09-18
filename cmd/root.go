package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	Version = "2.0"
)

var (
	Token string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lokalise",
	Short: "Lokalise CLI app",
	Long: `
=== Lokalise command-line tool ====

Lokalise is a CLI application that allows to perform queries to API. 
Enjoy!
`,
	Version: Version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// API Token, used for all commands
	rootCmd.PersistentFlags().StringVarP(&Token, "token", "t", "", "API token (required)")
	_ = rootCmd.MarkPersistentFlagRequired("token")

}
