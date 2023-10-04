## lokalise2 key list

List all keys

### Synopsis

Lists all keys in the project.

```
lokalise2 key list [flags]
```

### Options

```
      --disable-references uint8             Whether to disable key references.
      --filter-filenames string              One or more filenames to filter by.
      --filter-key-ids string                One or more key identifiers to filter by.
      --filter-keys string                   One or more key name to filter by. In case "Per-platform keys" is enabled in project settings, the filter will be applied to all platform names.
      --filter-platforms string              One or more platforms to filter by. Possible values are ios, android, web and other
      --filter-qa-issues string              One or more QA issues to filter by. Possible values are spelling_and_grammar, placeholders, html, url_count, url, email_count, email, brackets, numbers, leading_whitespace, trailing_whitespace, double_space and special_placeholder.
      --filter-tags string                   One or more tags to filter by.
      --filter-translation-lang-ids string   One or more language ID to filter by. Will include translations only for listed IDs.
      --filter-untranslated                  Filter by untranslated keys.
  -h, --help                                 help for list
      --include-comments uint8               Whether to include comments.
      --include-screenshots uint8            Whether to include URL to screenshots.
      --include-translations uint8           Whether to include translations.
```

### Options inherited from parent commands

```
      --config string       config file (default is ./config.yml)
      --project-id string   Unique project identifier (required).
  -t, --token string        API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 key](lokalise2_key.md)	 - Manage keys

