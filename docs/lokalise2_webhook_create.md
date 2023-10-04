## lokalise2 webhook create

Create a webhook

### Synopsis

Creates a webhook in the project. Requires `Manage settings` admin right.

```
lokalise2 webhook create [flags]
```

### Options

```
      --branch string           If webhook is limited to a single branch
      --event-lang-map string   Map the event with an array of languages iso codes. Omit this parameter for all languages in the project. JSON, see https://lokalise.com/api2docs/curl/#resource-webhooks
      --events strings          List of events to subscribe to (required, see https://developers.lokalise.com/docs/webhook-events).
  -h, --help                    help for create
      --url string              Specify the URL to your endpoint (required).
```

### Options inherited from parent commands

```
      --config string       config file (default is ./config.yml)
      --project-id string   Unique project identifier (required).
  -t, --token string        API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 webhook](lokalise2_webhook.md)	 - Manage webhooks

