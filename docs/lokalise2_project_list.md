## lokalise2 project list

List all projects

### Synopsis

Retrieves a list of projects available to the user, authorized with a token.

```
lokalise2 project list [flags]
```

### Options

```
      --filter-names string        One or more project names to filter by.
      --filter-team-id int         Limit results to team ID.
  -h, --help                       help for list
      --include-settings uint8     Whether to include project settings. (default 1)
      --include-statistics uint8   Whether to include project statistics. (default 1)
```

### Options inherited from parent commands

```
      --config string   config file (default is ./config.yml)
  -t, --token string    API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 project](lokalise2_project.md)	 - Manage projects

