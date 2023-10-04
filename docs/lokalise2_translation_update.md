## lokalise2 translation update

Update a translation

### Synopsis

Updates a translation.

```
lokalise2 translation update [flags]
```

### Options

```
      --custom-translation-status-ids strings   Custom translation status IDs to assign to translation (existing statuses will be replaced).
  -h, --help                                    help for update
      --is-fuzzy                                Whether the Fuzzy flag is enabled. (Note: Fuzzy is called Unverified in the editor now) .
      --is-reviewed                             Whether the Reviewed flag is enabled.
      --translation string                      The actual translation content. Use a JSON object for plural keys (required).
      --translation-id int                      A unique identifier of the translation (required).
```

### Options inherited from parent commands

```
      --config string       config file (default is ./config.yml)
      --project-id string   Unique project identifier (required).
  -t, --token string        API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 translation](lokalise2_translation.md)	 - Manage translations

