## lokalise2 task create

Create a task

### Synopsis

Creates a task in the project. Requires Manage tasks admin right.

```
lokalise2 task create [flags]
```

### Options

```
      --auto-close-languages                 Whether languages should be closed automatically upon completion of the last item (default true). Use --auto-close-languages=false to disable. (default true)
      --auto-close-task                      Whether the task should be automatically closed upon all language completion (default true). Use --auto-close-task=false to disable. (default true)
      --closing-tags strings                 Tags that will be added to affected keys when task is closed.
      --custom-translation-status-ids ints   IDs of custom translation statuses that will be applied to task items after item is completed.
      --description string                   Short description of the task.
      --do-lock-translations                 If set to 1, will lock translations for non-assigned project members.
      --due-date Y-m-d H:i:s                 Due date in Y-m-d H:i:s format. Example: `2018-12-24 23:59:59`.
  -h, --help                                 help for create
      --initial-tm-leverage                  Enable to calculate and save initial TM leverage with this task.
      --keys ints                            List of keys identifiers, included in task. Required if parent_task_id is not specified.
      --languages users                      List of languages in the task. One of users or `groups` must be provided (JSON, required, see https://lokalise.com/api2docs/curl/#transition-create-a-task-post).
      --parent-task-id int                   If task_type is review, it can have a parent task. Current task will be opened when parent task is closed.
      --task-type translation                Specify if task type is translation (default) or `review`.
      --title string                         Task title (required).
```

### Options inherited from parent commands

```
      --config string       config file (default is ./config.yml)
      --project-id string   Unique project identifier (required).
  -t, --token string        API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 task](lokalise2_task.md)	 - Manage tasks

