## lokalise2 key create

Create keys

### Synopsis

Creates one or more keys in the project.

```
lokalise2 key create [flags]
```

### Options

```
      --char-limit int             Maximum allowed number of characters in translations for this key.
      --comments strings           List of comments for this key.
      --context string             Optional context of the key (used with some file formats).
      --custom-attributes string   JSON containing custom attributes (if any).
      --description string         Description of the key.
      --filenames string           An object containing key filename attribute for each platform. (JSON, see https://lokalise.com/api2docs/curl/#transition-create-keys-post).
  -h, --help                       help for create
      --is-archived                Whether this key is archived.
      --is-hidden                  Whether this key is hidden from non-admins (translators).
      --is-plural                  Whether this key is plural.
      --key-name string            Key name. For projects with enabled Per-platform key names, pass JSON object with included ios, android, web and other string attributes. (JSON, required).
      --platforms strings          List of platforms, enabled for this key (required).
      --plural-name string         Optional custom plural name (used in some formats).
      --tags strings               List of tags for this keys.
      --translations string        Translations for all languages. (JSON, see https://lokalise.com/api2docs/curl/#transition-create-keys-post).
      --use-automations            Whether to run automations on the new key translations. (default true)
```

### Options inherited from parent commands

```
      --config string       config file (default is ./config.yml)
      --project-id string   Unique project identifier (required).
  -t, --token string        API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 key](lokalise2_key.md)	 - Manage keys

