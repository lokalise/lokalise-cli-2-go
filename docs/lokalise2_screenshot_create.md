## lokalise2 screenshot create

Create a screenshot

### Synopsis

Creates a screenshot in the project. Requires Manage screenshots admin right.

```
lokalise2 screenshot create [flags]
```

### Options

```
      --description string   Screenshot description.
      --file string          Path to a local image file (required).
  -h, --help                 help for create
      --key-ids uints        Attach the screenshot to key IDs specified. (default [])
      --ocr                  Try to recognize translations on the image and attach screenshot to all possible keys (default true). Use --ocr=false to disable. (default true)
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

