## lokalise2 webhook delete

Delete a webhook

### Synopsis

Deletes a configured webhook in the project. Requires Manage settings admin right.

```
lokalise2 webhook delete [flags]
```

### Options

```
  -h, --help                help for delete
      --webhook-id string   A unique identifier of the webhook (required).
```

### Options inherited from parent commands

```
      --config string       config file (default is ./config.yml)
      --project-id string   Unique project identifier (required).
  -t, --token string        API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 webhook](lokalise2_webhook.md)	 - Manage webhooks

