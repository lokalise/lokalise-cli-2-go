## lokalise2 contributor create

Create a contributor

### Synopsis

Creates a contributor in the project.
Requires Manage contributors admin right.

If is_admin flag is set to true, the user would automatically get access to all project languages, 
overriding supplied languages object. Attribute fullname will be ignored, 
if the user has already been registered in Lokalise.


```
lokalise2 contributor create [flags]
```

### Options

```
      --admin-rights strings   Custom list of user permissions. Possible values are activity, contributors, branches_create, branches_main_modify, branches_merge, custom_status_modify, download, glossary, glossary_edit, glossary_delete, keys, manage_languages, review, screenshots, settings, statistics, tasks, upload. Omitted or empty parameter will set no rights for the user.
      --email string           E-mail (required).
      --fullname string        Full name (only valid for inviting users, who previously did not have an account in Lokalise).
  -h, --help                   help for create
      --languages string       List of languages, accessible to the user. Required if is_admin is set to false (JSON, see https://lokalise.com/api2docs/curl/#transition-create-contributors-post).
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

