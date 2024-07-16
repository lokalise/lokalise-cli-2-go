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
      --admin-rights strings   Custom list of user permissions. Possible values are upload, activity, download, settings, statistics, keys, screenshots, contributors, languages. Omitted or empty parameter will set default admin rights for user role.
      --email string           E-mail (required).
      --fullname string        Full name (only valid for inviting users, who previously did not have an account in Lokalise).
  -h, --help                   help for create
      --is-admin               Whether the user has Admin access to the project.
      --is-reviewer            Whether the user has Reviewer access to the project.
      --languages string       List of languages, accessible to the user. Required if is_admin is set to false (JSON, see https://lokalise.com/api2docs/curl/#transition-create-contributors-post).
```

### Options inherited from parent commands

```
      --config string       config file (default is ./config.yml)
      --project-id string   Unique project identifier (required).
  -t, --token string        API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 contributor](lokalise2_contributor.md)	 - Manage project contributors

