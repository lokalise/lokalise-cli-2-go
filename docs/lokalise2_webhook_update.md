## lokalise2 webhook update

Update a webhook

### Synopsis

Updates a configured webhook in the project. Requires Manage settings admin right.

```
lokalise2 webhook update [flags]
```

### Options

```
      --branch string           If webhook is limited to a single branch
      --event-lang-map string   Map the event with an array of languages iso codes. Omit this parameter for all languages in the project.
      --events strings          Replace list of events, see https://developers.lokalise.com/docs/webhook-events.
  -h, --help                    help for update
      --url string              Update the URL to your endpoint.
```

### Options inherited from parent commands

```
      --config string       config file (default is ./config.yml)
      --project-id string   Unique project identifier (required).
  -t, --token string        API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 webhook](lokalise2_webhook.md)	 - Manage webhooks

