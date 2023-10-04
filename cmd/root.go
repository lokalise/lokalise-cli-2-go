package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/lokalise/go-lokalise-api/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	Version          = "2.6.11"
	DefaultPageLimit = 5000
)

var (
	Token string
	Api   *lokalise.Api
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "lokalise2",
	Short:   "Lokalise CLI v" + Version + ". Read the docs at https://github.com/lokalise/lokalise-cli-2-go",
	Version: Version,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		// init Api, runs like a middleware
		perPage := viper.GetUint("page-limit")
		if perPage == 0 {
			perPage = DefaultPageLimit
		}

		ClientOptions := []lokalise.ClientOption{
			lokalise.WithDebug(viper.GetBool("debug")),
			lokalise.WithRetryCount(viper.GetInt("retry-count")),
			lokalise.WithRetryTimeout(viper.GetDuration("retry-timeout")),
			lokalise.WithConnectionTimeout(viper.GetDuration("connection-timeout")),
			lokalise.WithPageLimit(perPage),
		}

		if viper.GetString("api-url") != "" {
			ClientOptions = append(ClientOptions, lokalise.WithBaseURL(viper.GetString("api-url")))
		}

		Api, err = lokalise.New(viper.GetString("token"), ClientOptions...)

		return err
	},
	DisableAutoGenTag: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if len(os.Args[1:]) > 0 && os.Args[1] == "gendocs" {
		fmt.Println("Generating docs...")
		err := doc.GenMarkdownTree(rootCmd, "./docs")
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(parseConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yml)")
	rootCmd.PersistentFlags().StringVarP(&Token, "token", "t", "", "API token. You can create API tokens at https://app.lokalise.com/profile.")

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

func printPageHeader(cur, total int64) {
	if viper.GetBool("debug") {
		fmt.Printf("\n=============\n Page %d of %d\n-------------\n", cur, total)
	}
}

// handy function for processing List response for all commands
func repeatableList(
	forwardPage func(page int64),
	list func() (lokalise.PageCounter, error),
) error {
	forwardPage(0)
	resp, err := list()
	if err != nil {
		return err
	}

	if resp.NumberOfPages() > 1 {
		printPageHeader(resp.CurrentPage(), resp.NumberOfPages())
		_ = printJson(resp)

		for p := resp.CurrentPage() + 1; p <= resp.NumberOfPages(); p++ {
			forwardPage(p)
			resp, err := list()
			if err != nil {
				return err
			}

			printPageHeader(p, resp.NumberOfPages())
			_ = printJson(resp)
		}
	} else {
		_ = printJson(resp)
	}

	return nil
}

// checkFlag checks if flag with a given name was among the ones being activated
func checkFlag(fs *pflag.FlagSet, name string) (wasSet bool) {
	fs.Visit(func(f *pflag.Flag) {
		if f.Name == name {
			wasSet = true
		}
	})

	return
}
