## lokalise2 translation list

List all translations

### Synopsis

Retrieves a list of project translation items, ungrouped. You may want to request Keys resource in order to get the structured key/translation pairs for all languages.

```
lokalise2 translation list [flags]
```

### Options

```
      --disable-references uint8   Whether to disable key references.
      --filter-fuzzy uint8         Filter translations which are unverified (fuzzy).
      --filter-is-reviewed uint8   Filter translations which are reviewed.
      --filter-lang-id string      Return translations only for presented language ID.
      --filter-qa-issues string    One or more QA issues to filter by. Possible values are spelling_and_grammar, placeholders, html, url_count, url, email_count, email, brackets, numbers, leading_whitespace, trailing_whitespace, double_space and special_placeholder.
  -h, --help                       help for list
```

### Options inherited from parent commands

```
      --config string       config file (default is ./config.yml)
      --project-id string   Unique project identifier (required).
  -t, --token string        API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 translation](lokalise2_translation.md)	 - Manage translations

