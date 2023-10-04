## lokalise2 language

Manage languages

### Synopsis

There are over 500 predefined language/dialect combinations available in Lokalise. In case you require a custom language/dialect combination use custom_X languages (where X is a number from 1 to 100). You may override language code and name when adding a language, or update an existing language properties later.

There are situations when it is necessary to export different language codes to different platforms (e.g. zh-Hans to iOS and zh_Hans to Web). In such cases you need to set any preferred version and ise export parameter to set language mapping depending on the file format.


### Options

```
  -h, --help                help for language
      --project-id string   Unique project identifier (required).
```

### Options inherited from parent commands

```
      --config string   config file (default is ./config.yml)
  -t, --token string    API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2](lokalise2.md)	 - Lokalise CLI v2.6.11. Read the docs at https://github.com/lokalise/lokalise-cli-2-go
* [lokalise2 language create](lokalise2_language_create.md)	 - Create languages
* [lokalise2 language delete](lokalise2_language_delete.md)	 - Delete a language
* [lokalise2 language list](lokalise2_language_list.md)	 - List project languages
* [lokalise2 language list-system](lokalise2_language_list-system.md)	 - List system languages
* [lokalise2 language retrieve](lokalise2_language_retrieve.md)	 - Retrieve a language
* [lokalise2 language update](lokalise2_language_update.md)	 - Update a language

