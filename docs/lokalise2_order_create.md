## lokalise2 order create

Create an order

### Synopsis

Creates a translation order. You must have admin privileges in the project you are placing an order in.

```
lokalise2 order create [flags]
```

### Options

```
      --briefing string                Order briefing (required).
      --card-id int                    Card identifier that should be used for payment. (required).
      --dry-run                        Return the response without actually placing an order. Useful for price estimation. The card will not be charged.
  -h, --help                           help for create
      --keys ints                      List of keys identifiers, included in the order (required).
      --project-id string              Project identifier. (required).
      --provider-slug string           Translation provider slug (required).
      --source-language-iso string     Source language code of the order (required).
      --target-language-isos strings   List of target languages (required).
      --translation-style string       Only for gengo provider. Available values are formal, informal, business, friendly. Defaults to friendly.
      --translation-tier int           Tier of the translation. Tiers depend on provider (order).
```

### Options inherited from parent commands

```
      --config string   config file (default is ./config.yml)
      --team-id int     A unique identifier of the team (required).
  -t, --token string    API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 order](lokalise2_order.md)	 - Manage orders

