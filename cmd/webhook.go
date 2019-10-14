package cmd

import (
	"encoding/json"
	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

var (
	webhookId string

	newWebhook       lokalise.CreateWebhook
	updateWebhook    lokalise.UpdateWebhook
	eventLanguageMap string
)

// snapshotCmd represents the snapshot command
var webhookCmd = &cobra.Command{
	Use:   "webhook",
	Short: "Manage webhooks",
}

var webhookListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all webhooks",
	Long:  "Retrieves a list of configured webhooks. Requires `Manage settings` admin right.",
	RunE: func(*cobra.Command, []string) error {
		c := Api.Webhooks()
		pageOpts := c.PageOpts()

		return repeatableList(
			func(p int64) {
				pageOpts.Page = uint(p)
				c.SetPageOptions(pageOpts)
			},
			func() (lokalise.PageCounter, error) {
				return c.List(projectId)
			},
		)
	},
}

var webhookRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieve a webhook",
	Long:  "Retrieves a Webhook object. Requires Manage settings admin right.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Webhooks().Retrieve(projectId, webhookId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var webhookCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a webhook",
	Long:  "Creates a webhook in the project. Requires `Manage settings` admin right.",
	RunE: func(*cobra.Command, []string) error {
		// preparing params
		if eventLanguageMap != "" {
			err := json.Unmarshal([]byte(eventLanguageMap), &newWebhook.EventLangMap)
			if err != nil {
				return err
			}
		}

		resp, err := Api.Webhooks().Create(projectId, newWebhook)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var webhookUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a webhook",
	RunE: func(*cobra.Command, []string) error {
		// preparing params
		if eventLanguageMap != "" {
			err := json.Unmarshal([]byte(eventLanguageMap), &updateWebhook.EventLangMap)
			if err != nil {
				return err
			}
		}

		resp, err := Api.Webhooks().Update(projectId, webhookId, updateWebhook)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var webhookDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a webhook",
	Long:  "Deletes a configured webhook in the project. Requires Manage settings admin right.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Webhooks().Delete(projectId, webhookId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	webhookCmd.AddCommand(webhookListCmd, webhookCreateCmd, webhookRetrieveCmd, webhookUpdateCmd, webhookDeleteCmd)
	rootCmd.AddCommand(webhookCmd)

	// general flags
	flagProjectId(webhookCmd, true)

	// Create
	fs := webhookCreateCmd.Flags()
	fs.StringVar(&newWebhook.URL, "url", "", "Specify the URL to your endpoint (required).")
	_ = webhookCreateCmd.MarkFlagRequired("url")
	fs.StringSliceVar(&newWebhook.Events, "events", []string{}, "List of events to subscribe to (required).")
	_ = webhookCreateCmd.MarkFlagRequired("events")
	fs.StringVar(&eventLanguageMap, "event-lang-map", "", "Map the event with an array of languages iso codes. Omit this parameter for all languages in the project")

	// Update
	fs = webhookUpdateCmd.Flags()
	fs.StringVar(&updateWebhook.URL, "url", "", "Update the URL to your endpoint.")
	fs.StringSliceVar(&updateWebhook.Events, "events", []string{}, "Replace list of events, see https://docs.lokalise.com/en/articles/3184756-webhooks for docs")
	fs.StringVar(&eventLanguageMap, "event-lang-map", "", "Map the event with an array of languages iso codes. Omit this parameter for all languages in the project")

	// Retrieve, delete
	flagWebhookId(webhookRetrieveCmd)
	flagWebhookId(webhookDeleteCmd)
}

func flagWebhookId(cmd *cobra.Command) {
	cmd.Flags().StringVar(&webhookId, "webhook-id", "", "A unique identifier of the webhook (required).")
	_ = cmd.MarkFlagRequired("webhook-id")
}
