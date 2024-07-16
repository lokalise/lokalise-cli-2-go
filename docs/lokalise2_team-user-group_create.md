## lokalise2 team-user-group create

Create a new user group

### Synopsis

Creates a group in the team. Requires Admin right in the team.

```
lokalise2 team-user-group create [flags]
```

### Options

```
      --admin-rights strings   List of group administrative permissions. Required if group has admin rights.
  -h, --help                   help for create
      --is-admin               Whether the group has Admin access to the project (required).
      --is-reviewer            Whether the group has reviewer access to the project (required).
      --languages string       List of languages. Required if group doesn't have admin rights. JSON, see https://lokalise.com/api2docs/curl/#transition-create-a-group-post
      --name string            Name of the group (required).
```

### Options inherited from parent commands

```
      --config string   config file (default is ./config.yml)
      --team-id int     A unique identifier of the team (required).
  -t, --token string    API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 team-user-group](lokalise2_team-user-group.md)	 - Manage team user groups

