## lokalise2 team-user update

Update a team user

### Synopsis

Updates the role of a team user. Requires Admin role in the team.

```
lokalise2 team-user update [flags]
```

### Options

```
  -h, --help          help for update
      --role string   Role of the user. Available roles are owner, admin, member (required).
      --user-id int   A unique identifier of the user (required).
```

### Options inherited from parent commands

```
      --config string   config file (default is ./config.yml)
      --team-id int     A unique identifier of the team (required).
  -t, --token string    API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 team-user](lokalise2_team-user.md)	 - Manage team users

