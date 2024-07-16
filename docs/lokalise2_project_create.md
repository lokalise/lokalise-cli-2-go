## lokalise2 project create

Create a project

### Synopsis

Creates a new project in the specified team. Requires Admin role in the team.

```
lokalise2 project create [flags]
```

### Options

```
      --base-lang-iso string   Language/locale code of the project base language. Should be in a scope of languages list. Use custom_iso code in case it was defined.
      --description string     Description of the project.
  -h, --help                   help for create
      --languages string       List of languages to add (JSON, see https://lokalise.com/api2docs/curl/#transition-create-a-project-post).
      --name string            Name of the project (required).
      --project-type string    Project type. Allowed values are localization_files, paged_documents.
      --team-id int            ID of the team to create a project in. If this parameter is omitted, the project will be created in current team of the user, whose API token is specified.
```

### Options inherited from parent commands

```
      --config string   config file (default is ./config.yml)
  -t, --token string    API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 project](lokalise2_project.md)	 - Manage projects

