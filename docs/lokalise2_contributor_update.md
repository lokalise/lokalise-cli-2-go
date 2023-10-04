## lokalise2 contributor update

Update a contributor

### Synopsis

Updates a contributor.
Requires Manage contributors admin right.

If you want to give an existing contributor access to a new language, you must specify full languages array, 
including the previously added languages as well.


```
lokalise2 contributor update [flags]
```

### Options

```
      --admin-rights strings   Custom list of user permissions. Possible values are upload, activity, download, settings, statistics, keys, screenshots, contributors, languages. Empty parameter will set default admin rights for user role.
      --contributor-id int     A unique identifier of contributor (required).
  -h, --help                   help for update
      --is-admin               Whether the user has Admin access to the project.
      --is-reviewer            Whether the user has Reviewer access to the project.
      --languages string       List of languages, accessible to the user (JSON, see https://lokalise.com/api2docs/curl/#transition-update-a-contributor-put).
```

### Options inherited from parent commands

```
      --config string       config file (default is ./config.yml)
      --project-id string   Unique project identifier (required).
  -t, --token string        API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 contributor](lokalise2_contributor.md)	 - Manage project contributors

