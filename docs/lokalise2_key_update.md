## lokalise2 key update

Update a key

### Synopsis

Updates the properties of a key and it’s associated objects. Requires Manage keys admin right.

```
lokalise2 key update [flags]
```

### Options

```
      --char-limit int             Maximum allowed number of characters in translations for this key.
      --context string             Optional context of the key (used with some file formats).
      --custom-attributes string   JSON containing custom attributes (if any).
      --description string         Description of the key.
      --filenames string           An object containing key filename attribute for each platform. (JSON, see https://lokalise.com/api2docs/curl/#transition-update-a-key-put).
  -h, --help                       help for update
      --is-archived                Whether this key is archived.
      --is-hidden                  Whether this key is hidden from non-admins (translators).
      --is-plural                  Whether this key is plural.
      --key-id int                 A unique identifier of the key (required).
      --key-name object            Key identifier. For projects with enabled Per-platform key names, pass object with included ios, android, web and other string attributes.
      --merge-tags                 Enable to merge specified tags with the current tags attached to the key.
      --platforms strings          List of platforms, enabled for this key. Possible values are ios, android, web and other.
      --plural-name string         Optional custom plural name (used in some formats).
      --tags strings               List of tags for this keys.
```

### Options inherited from parent commands

```
      --config string       config file (default is ./config.yml)
      --project-id string   Unique project identifier (required).
  -t, --token string        API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 key](lokalise2_key.md)	 - Manage keys

