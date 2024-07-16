## lokalise2 task update

Update a task

### Synopsis

Updates the properties of a task. Requires Manage tasks admin right.

```
lokalise2 task update [flags]
```

### Options

```
      --auto-close-languages   Whether languages should be closed automatically upon completion of the last item (default true). Use --auto-close-languages=false to disable. (default true)
      --auto-close-task        Whether the task should be automatically closed upon all language completion (default true). Use --auto-close-task=false to disable. (default true)
      --close-task             Whether the task should be closed and notifications sent. The task cannot be reopened again.
      --closing-tags strings   Tags that will be added to affected keys when task is closed.
      --description string     Short description of the task.
      --do-lock-translations   If set to 1, will lock translations for non-assigned project members.
      --due-date Y-m-d H:i:s   Due date in Y-m-d H:i:s format. Example: `2018-12-24 23:59:59`.
  -h, --help                   help for update
      --languages string       List of languages to update. JSON, differs from a creation list, see https://lokalise.com/api2docs/curl/#transition-update-a-task-put.
      --task-id int            A unique identifier of the task (required).
      --title string           Task title.
```

### Options inherited from parent commands

```
      --config string       config file (default is ./config.yml)
      --project-id string   Unique project identifier (required).
  -t, --token string        API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 task](lokalise2_task.md)	 - Manage tasks

