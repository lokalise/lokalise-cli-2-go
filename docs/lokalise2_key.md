## lokalise2 key

Manage keys

### Synopsis

Keys are core item elements of the project.

Each phrase that is used in your app or website must be identified by a key and have values that represent translations to various languages. For example key index.welcome would have values of Welcome in English and Benvenuto in Italian. Keys can be assigned to one or multiple platforms. Once a key is assigned to a platform, it would be included in the export for file formats related to this platform.


### Options

```
  -h, --help                help for key
      --project-id string   Unique project identifier (required).
```

### Options inherited from parent commands

```
      --config string   config file (default is ./config.yml)
  -t, --token string    API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2](lokalise2.md)	 - Lokalise CLI v2.6.11. Read the docs at https://github.com/lokalise/lokalise-cli-2-go
* [lokalise2 key create](lokalise2_key_create.md)	 - Create keys
* [lokalise2 key delete](lokalise2_key_delete.md)	 - Delete a key
* [lokalise2 key list](lokalise2_key_list.md)	 - List all keys
* [lokalise2 key retrieve](lokalise2_key_retrieve.md)	 - Retrieve a key
* [lokalise2 key update](lokalise2_key_update.md)	 - Update a key

