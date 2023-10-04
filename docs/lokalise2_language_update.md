## lokalise2 language update

Update a language

### Synopsis

Updates the properties of a language. Requires Manage languages admin right.

```
lokalise2 language update [flags]
```

### Options

```
  -h, --help                   help for update
      --lang-id int            A unique identifier of the language (required).
      --lang-iso string        Language/locale code.
      --lang-name string       Language name.
      --plural-forms strings   List of supported plural forms.
```

### Options inherited from parent commands

```
      --config string       config file (default is ./config.yml)
      --project-id string   Unique project identifier (required).
  -t, --token string        API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 language](lokalise2_language.md)	 - Manage languages

