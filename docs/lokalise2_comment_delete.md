## lokalise2 comment delete

Delete a comment

### Synopsis

Deletes a comment from the project. Authenticated user can only delete own comments.

```
lokalise2 comment delete [flags]
```

### Options

```
      --comment-id int   A unique identifier of comment (required).
  -h, --help             help for delete
      --key-id int       A unique identifier of the key (required).
```

### Options inherited from parent commands

```
      --config string       config file (default is ./config.yml)
      --project-id string   Unique project identifier (required).
  -t, --token string        API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 comment](lokalise2_comment.md)	 - Manage key comments

