## lokalise2 file

Upload and download files

### Synopsis

Lokalise is a project-oriented translation management system, which means we store all keys and translations in the database and can generate files in any format you require. Assigning a key to one or more platforms means including the key in the export routine for file formats, associated with this platform, e.g. if a key is assigned to iOS platform it would get included with strings and xliff format exports. In addition to assign keys to platforms you may assign keys to files and have different filename depending on the platform. List of supported file formats is available here https://docs.lokalise.com/en/collections/2909121-keys-and-files#supported-file-formats.

### Options

```
  -h, --help                help for file
      --project-id string   Unique project identifier (required).
```

### Options inherited from parent commands

```
      --config string   config file (default is ./config.yml)
  -t, --token string    API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2](lokalise2.md)	 - Lokalise CLI v2.6.11. Read the docs at https://github.com/lokalise/lokalise-cli-2-go
* [lokalise2 file download](lokalise2_file_download.md)	 - Download files
* [lokalise2 file list](lokalise2_file_list.md)	 - List all files
* [lokalise2 file upload](lokalise2_file_upload.md)	 - Upload a file

