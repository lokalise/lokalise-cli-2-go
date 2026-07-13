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
      --admin-rights strings   Custom list of user permissions. Possible values are activity, contributors, branches_create, branches_main_modify, branches_merge, custom_status_modify, download, glossary, glossary_edit, glossary_delete, keys, manage_languages, review, screenshots, settings, statistics, tasks, upload. Empty parameter will set no rights for the user.
      --contributor-id int     A unique identifier of contributor (required).
  -h, --help                   help for update
      --languages string       List of languages, accessible to the user (JSON, see https://lokalise.com/api2docs/curl/#transition-update-a-contributor-put).
      --role-id int            Permission template id for the contributor. By setting this admin_rights will be ignored and a template will be assigned with predefined permission set.
```

### Options inherited from parent commands

```
      --config string       config file (default is ./config.yml)
      --project-id string   Unique project identifier (required).
  -t, --token string        API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 contributor](lokalise2_contributor.md)	 - Manage project contributors

