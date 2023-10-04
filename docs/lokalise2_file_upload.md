## lokalise2 file upload

Upload a file

### Synopsis

Imports a localization file to the project. Requires Upload files admin right. List of supported file formats is available here https://docs.lokalise.com/en/collections/2909121-keys-and-files#supported-file-formats.

```
lokalise2 file upload [flags]
```

### Options

```
      --apply-tm                                   Enable to automatically apply 100% translation memory matches.
      --cleanup-mode                               Enable to delete all keys with all language translations that are not present in the uploaded file. You may want to make a snapshot of the project before importing new file, just in case.
      --convert-placeholders                       Enable to automatically convert placeholders to the Lokalise universal placeholders. (default true)
      --custom-translation-status-ids int64Slice   Custom translation status IDs to be added to translations. By default statuses are applied to created and updated translations. (default [])
      --custom-translation-status-inserted-keys    Add specified custom translation statuses to inserted keys (default true). Use --custom-translation-status-inserted-keys=false to disable. (default true)
      --custom-translation-status-skipped-keys     Add specified custom translation statuses to skipped keys.
      --custom-translation-status-updated-keys     Add specified custom translation statuses to updated keys (default true). Use --custom-translation-status-updated-keys=false to disable. (default true)
      --detect-icu-plurals                         Enable to automatically detect and parse ICU formatted plurals in your translations.
      --distinguish-by-file                        Enable to allow keys with similar names to coexist, in case they are assigned to different filenames.
      --file string                                Path to local file (required).
  -h, --help                                       help for upload
      --hidden-from-contributors                   Enable to automatically set newly created keys as 'Hidden from contributors'
      --include-path                               Include relative directory name in the filename when uploading.
      --keys-to-values                             Enable to automatically replace values with key names.
      --lang-iso string                            Language code of the translations in the file you are importing (required).
      --poll                                       Enable to wait until background file upload finishes with result
      --poll-timeout duration                      Specify custom file upload polling maximum duration. Default: 30s (default 30s)
      --replace-modified                           Enable to replace translations, that have been modified (in the file being uploaded).
      --skip-detect-lang-iso                       Skip automatic language detection by filename. Default: false
      --slashn-to-linebreak                        Enable to replace \n with a line break (default true). Use --slashn-to-linebreak=false to disable. (default true)
      --tag-inserted-keys                          Add specified tags to inserted keys (default true). Use --tag-inserted-keys=false to disable (default true)
      --tag-skipped-keys                           Add specified tags to skipped keys.
      --tag-updated-keys                           Add specified tags to updated keys (default true). Use tag-updated-keys=false to disable (default true)
      --tags strings                               Tag keys with the specified tags. By default tags are applied to created and updated keys.
      --use-automations                            Whether to run automations for this upload. (default true)
```

### Options inherited from parent commands

```
      --config string       config file (default is ./config.yml)
      --project-id string   Unique project identifier (required).
  -t, --token string        API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 file](lokalise2_file.md)	 - Upload and download files

