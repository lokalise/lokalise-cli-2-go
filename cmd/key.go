package cmd

import (
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

// keyCmd represents the key command
var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Keys are core item elements of the project.",
	Long: `
Each phrase that is used in your app or website must be identified by a key and have values 
that represent translations to various languages. For example key 'index.welcome' would have values of
'Welcome' in English and 'Benvenuto' in Italian. 
Keys can be assigned to one or multiple platforms. Once a key is assigned to a platform, 
it would be included in the export for file formats related to this platform.

One of the unique features of Lokalise is the ability to use similar keys across different platforms, 
thus reducing the translation work amount to be done by translators. 
Once you import or add keys, they need to be assigned to one or several platforms (e.g. iOS, Android).

See also https://docs.lokalise.com/developer-docs/keys-and-platforms
`,
}

var keyListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all keys",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Keys().List(projectId, lokalise.ListKeysOptions{}) // todo check
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var keyCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates one or more keys in the project.",
	/*RunE: func(cmd *cobra.Command, args []string) error {
		c := lokalise.Key{key} // todo implement multiple keys
		resp, err := Api.Keys().Create(projectId, []lokalise.Key{c})
		if err != nil {
			return err
		}
		return printJson(resp)
	},*/
}

var keyRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a Key object",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Keys().Retrieve(projectId, keyId, lokalise.RetrieveKeyOptions{})
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var keyUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a Key object",
	/*RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Keys().Update(projectId, keyId, key) // todo also bulk update here for key-names as json
		if err != nil {
			return err
		}
		return printJson(resp)
	},*/
}

var keyDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a key from the project.",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := Api.Keys().Delete(projectId, keyId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	keyCmd.AddCommand(keyListCmd)
	keyCmd.AddCommand(keyCreateCmd)
	keyCmd.AddCommand(keyRetrieveCmd)
	keyCmd.AddCommand(keyUpdateCmd)
	keyCmd.AddCommand(keyDeleteCmd)

	rootCmd.AddCommand(keyCmd)

	// common for all Comment cmd`s
	keyCmd.PersistentFlags().StringVar(&projectId, "project-id", "", "A unique project identifier (required)")
	_ = keyCmd.MarkPersistentFlagRequired("project-id")

	// separate flags for every command
	withKeyId(keyCreateCmd)
	withKeyId(keyRetrieveCmd)
	withKeyId(keyDeleteCmd)
}

func withKeyId(cmd *cobra.Command) { // todo rename as flagKeyId
	cmd.Flags().Int64Var(&keyId, "key-id", 0, "A unique identifier of key (required)")
	_ = cmd.MarkFlagRequired("key-id")
}
