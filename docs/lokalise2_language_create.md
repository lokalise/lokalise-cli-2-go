## lokalise2 language create

Create languages

### Synopsis

Creates one or more languages in the project. Requires Manage languages admin right.

	The lang-iso is the identifer of one of the system languages. You are only required to include the lang-iso attribute, however you may override the default language code, language name and plural forms as well.


```
lokalise2 language create [flags]
```

### Options

```
      --custom-iso string             Override language/locale code.
      --custom-name string            Override language name.
      --custom-plural-forms strings   Override list of supported plural forms for this language.
  -h, --help                          help for create
      --lang-iso string               A unique language code in the system.
```

### Options inherited from parent commands

```
      --config string       config file (default is ./config.yml)
      --project-id string   Unique project identifier (required).
  -t, --token string        API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 language](lokalise2_language.md)	 - Manage languages

