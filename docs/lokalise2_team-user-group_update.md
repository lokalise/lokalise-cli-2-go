## lokalise2 team-user-group update

Update a user group

### Synopsis

Updates the properties of a group. Requires Admin right in the team

```
lokalise2 team-user-group update [flags]
```

### Options

```
      --admin-rights strings   List of group permissions.
      --group-id int           A unique identifier of the group (required).
  -h, --help                   help for update
      --is-admin               Whether the group has Admin access to the project. Deprecated and will be removed, but still required as of now.
      --languages string       List of languages. Required if group doesn't have admin rights.
      --name string            Name of the group (required).
      --role-id int            Permission template id for the contributor. By setting this admin-rights will be ignored and a template will be assigned with predefined permission set.
```

### Options inherited from parent commands

```
      --config string   config file (default is ./config.yml)
      --team-id int     A unique identifier of the team (required).
  -t, --token string    API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 team-user-group](lokalise2_team-user-group.md)	 - Manage team user groups

