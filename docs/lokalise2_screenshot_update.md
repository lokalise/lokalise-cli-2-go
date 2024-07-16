## lokalise2 screenshot update

Updates a screenshot

### Synopsis

Updates properties of a screenshot. Requires Manage screenshots admin right.

```
lokalise2 screenshot update [flags]
```

### Options

```
      --description string   Screenshot description.
  -h, --help                 help for update
      --key-ids uints        Attach the screenshot to key IDs specified. (default [])
      --screenshot-id int    A unique identifier of the screenshot (required).
      --tags strings         List of tags to add to the uploaded screenshot.
      --title string         Screenshot title
```

### Options inherited from parent commands

```
      --config string       config file (default is ./config.yml)
      --project-id string   Unique project identifier (required).
  -t, --token string        API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 screenshot](lokalise2_screenshot.md)	 - Manage screenshots

