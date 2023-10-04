## lokalise2 snapshot restore

Restore a snapshot

### Synopsis

Restores project snapshot to a project copy. Requires Manage settings admin right and Admin role in the team.

```
lokalise2 snapshot restore [flags]
```

### Options

```
  -h, --help              help for restore
      --snapshot-id int   A unique identifier of the snapshot (required).
```

### Options inherited from parent commands

```
      --config string       config file (default is ./config.yml)
      --project-id string   Unique project identifier (required).
  -t, --token string        API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 snapshot](lokalise2_snapshot.md)	 - Manage snapshots

