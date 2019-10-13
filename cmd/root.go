package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	Version          = "2.0"
	DefaultPageLimit = 5000
)

var (
	Token string
	Api   *lokalise.Api
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "lokalise2",
	Short:   "Lokalise CLI v" + Version + ". Documentation is available at https://docs.lokalise.com/cli2",
	Version: Version,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		// init Api, runs like a middleware
		perPage := viper.GetUint("page-limit")
		if perPage == 0 {
			perPage = DefaultPageLimit
		}

		Api, err = lokalise.New(
			viper.GetString("token"),

			lokalise.WithDebug(viper.GetBool("debug")),
			lokalise.WithRetryCount(viper.GetInt("retry-count")),
			lokalise.WithRetryTimeout(viper.GetDuration("retry-timeout")),
			lokalise.WithConnectionTimeout(viper.GetDuration("connection-timeout")),
			lokalise.WithPageLimit(perPage),
		)
		return err
	},
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
	// init API Token, used for all commands
	if viper.GetString("token") == "" {
		// if not found in config
		rootCmd.PersistentFlags().StringVarP(&Token, "token", "t", "", "API token (required). You can create API tokens at https://lokalise.com/profile.")
		_ = rootCmd.MarkPersistentFlagRequired("token")

	} else {
		rootCmd.PersistentFlags().StringVarP(&Token, "token", "t", "", "API token (override value from config if desired).")
	}
	// binding
	_ = viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
}

// ‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Utils
// _____________________________________________________________________________________________________________________

func printJson(v interface{}) error {
	output, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(output))
	return nil
}
