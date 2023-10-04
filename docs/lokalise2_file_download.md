## lokalise2 file download

Download files

### Synopsis

Exports project files as a .zip bundle. Generated bundle will be uploaded to an Amazon S3 bucket, which will be stored there for 12 months available to download. As the bundle is generated and uploaded you would get a response with the URL to the file. Requires Download files admin right.

```
lokalise2 file download [flags]
```

### Options

```
      --add-newline-eof                         Enable to add new line at end of file (if supported by format).
      --all-platforms                           Enable to include all platform keys. If disabled, only the keys, associated with the platform of the format will be exported.
      --bundle-description string               Description of the created bundle. Applies to ios_sdk or android_sdk OTA SDK bundles.
      --bundle-structure string                 Bundle structure, used when original-filenames set to false. Allowed placeholders are %LANG_ISO%, %LANG_NAME%, %FORMAT% and %PROJECT_NAME%).
      --custom-translation-status-ids strings   Only translations attributed to selected custom statuses will be included. Leave empty for all.
      --dest string                             Destination folder for ZIP file. (default "./")
      --directory-prefix string                 Directory prefix in the bundle (used when original-filenames set to true). Allowed placeholder is %LANG_ISO%.
      --disable-references                      Enable to skip automatic replace of key reference placeholders (e.g. [%key:hello_world%]) with their corresponding translations. In case you have this disabled and are still getting references, make sure the permissions of the projects are configured right.
      --escape-percent                          Only works for printf placeholder format. When enabled, all universal percent placeholders "[%]" will be always exported as "%%".
      --exclude-tags strings                    Specify to exclude keys with these tags.
      --export-empty-as string                  Select how you would like empty translations to be exported. Allowed values are empty to keep empty, base to replace with the base language value, or skip to omit.
      --export-sort string                      Export key sort mode. Allowed value are first_added, last_added, last_updated, a_z, z_a.
      --filter-data strings                     Narrow export data range. Allowed values are translated or untranslated, reviewed (or reviewed_only), last_reviewed_only, nonfuzzy and nonhidden. (Note: Fuzzy is called Unverified in the editor now).
      --filter-filenames strings                Only keys attributed to selected files will be included. Leave empty for all.
      --filter-langs strings                    List of languages to export. Omit this parameter for all languages.
      --filter-repositories strings             Pull requests will be created only for listed repositories (organization/repository format). Leave empty array to process all configured integrations by platform only.
      --format string                           File format (e.g. json, strings, xml). Must be file extension of any of the file formats we support. May also be ios_sdk or android_sdk for respective OTA SDK bundles. (required)
  -h, --help                                    help for download
      --icu-numeric                             If enabled, plural forms zero, one and two will be replaced with =0, =1 and =2 respectively. Only works for ICU plural format.
      --include-comments                        Enable to include key comments and description in exported file (if supported by the format).
      --include-description                     Enable to include key description in exported file (if supported by the format) (default true). Use --include-description=false to disable. (default true)
      --include-pids strings                    Other projects ID's, which keys should be included with this export.
      --include-tags strings                    Narrow export range to tags specified.
      --indentation string                      Provide to override default indentation in supported files. Allowed values are default, 1sp, 2sp, 3sp, 4sp, 5sp, 6sp, 7sp, 8sp and tab.
      --java-properties-encoding string         (Java Properties export only). Encoding for .properties files. Allowed values are utf-8 and latin-1.
      --java-properties-separator string        (Java Properties export only). Separator for keys/values in .properties files. Allowed values are = and :.
      --json-only                               Should only the API JSON response be returned.
      --json-unescaped-slashes                  (JSON export only). Enable to leave forward slashes unescaped.
      --keep-zip                                Keep or delete ZIP file after being unpacked.
      --language-mapping string                 List of languages to override default iso codes for this export (JSON, see https://lokalise.com/api2docs/curl/#transition-download-files-post).
      --original-filenames                      Enable to use original filenames/formats. If set to false (--original-filenames=false) all keys will be export to a single file per language (default true). (default true)
      --placeholder-format string               Override the default placeholder format for the file type. Allowed values are printf, ios, icu, net, symfony, i18n, raw.
      --plural-format string                    Override the default plural format for the file type. Allowed values are json_string, icu, array, generic, symfony, i18next.
      --replace-breaks                          Enable to replace line breaks in exported translations with \n (default true). Use --replace-breaks=false to disable. (default true)
      --triggers strings                        Trigger integration exports (must be enabled in project settings). Allowed values are amazons3, gcs, github, github-enterprise, gitlab, bitbucket, bitbucket-enterprise.
      --unzip-to string                         Unzip to this folder. (default "./")
      --webhook-url string                      Once the export is complete, sends a HTTP POST with the generated bundle URL to the specified URL.
      --yaml-include-root                       (YAML export only). Enable to include language ISO code as root key.
```

### Options inherited from parent commands

```
      --config string       config file (default is ./config.yml)
      --project-id string   Unique project identifier (required).
  -t, --token string        API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 file](lokalise2_file.md)	 - Upload and download files

