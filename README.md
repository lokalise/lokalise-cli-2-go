# lokalise-cli-2-go
Lokalise CLI 2.0 Golang sources

Objectives:

- Develop a CLI tool for Lokalise (https://lokalise.com) API 2.0 (https://lokalise.com/api2docs/curl/#resource-getting-started). 
- The CLI tool must be developed using GO programming language using either of https://github.com/urfave/cli or https://github.com/spf13/cobra. 
- The resulting app must be compiled for linux, mac, windows.
- All the output should be mirrored from API as JSON (besides the `file download` command, where there is a switch to show json only, otherwise download the file too). 
- The CLI tool must respect pagination and sequentally fetch all pages if more than 1 returned by API.

Command line tool parameter specifications:

lokalise

	--config string

	--version

	--token string*

	--no-pretty bool

	--cli-read-timeout int

	--cli-connect-timeout int

	comment
		list
			--project-id string*

		list-key
			--project-id string*
			--key-id int*

		create
			--project-id string*
			--key-id int*
			--comment int*

		retrieve
			--project-id string*
			--key-id int*
			--comment-id int*

		delete
			--project-id string*
			--key-id int*
			--comment-id int*

	contributor
		list
			--project-id string*

		create
			--project-id string*
			--email string*
			--fullname string
			--is-admin bool
			--is-reviewer bool
			--languages json
			--admin-rights array

		retrieve
			--project-id string*
			--contributor-id int*

		update
			--project-id string*
			--contributor-id int*
			--is-admin bool
			--is-reviewer bool
			--languages json
			--admin-rights array						

		delete
			--project-id string*
			--contributor-id int*	

	file
		list
			--project-id string*
			--filter-filename string

		upload
			--project-id string*
			--file string* (+masks)
			--force-filename (string, only possible if a single file is supplied)
			--lang-iso string*
			--convert-placeholders bool
			--detect-icu-plurals bool
			--tags array
			--tag-inserted-keys bool
			--tag-updated-keys bool
			--tag-skipped-keys bool
			--replace-modified bool
			--slashn-to-linebreak bool
			--keys-to-values bool
			--distinguish-by-file bool
			--apply-tm bool
			--hidden-from-contributors bool
			--cleanup-mode bool
		
		download
			--project-id string*
			--format string*
			--json-only bool (default false, triggers if only the API JSON response should be returned. If false, also download the file)
			--dest string (destination folder for zip file)
			--unzip-to string (unzip to this folder)
			--keep-zip bool (default false, keep or delete zip after unzip)
			--original-filenames bool
			--bundle-structure string
			--directory-prefix string
			--all-platforms bool
			--filter-langs array
			--filter-data array
			--filter-filenames array
			--add-newline-eof bool
			--custom-translation-status-ids array
			--include-tags array
			--exclude-tags array
			--export-sort string
			--export-empty-as string
			--include-comments bool
			--include-description bool
			--include-pids array
			--triggers array
			--filter-repositories array
			--replace-breaks bool
			--disable-references bool
			--plural-format string
			--placeholder-format string
			--webhook-url string
			--language-mapping json
			--icu-numeric bool
			--escape-percent 
			--indentation string
			--yaml-include-root bool
			--json-unescaped-slashes bool
			--java-properties-encoding string
			--java-properties-separator string
			--bundle-description string

	key
		list
			--project-id string*
			--disable-references bool
			--include-comments bool
			--include-screenshots bool
			--include-translations bool
			--filter-translation-lang-ids array
			--filter-tags array
			--filter-filenames array
			--filter-keys array
			--filter-key-ids array
			--filter-platforms array (Possible values are ios, android, web and other)
			--filter-untranslated bool
			--filter-qa-issues array (Possible values are spelling-and-grammar, placeholders, html, url-count, url, email-count, email, brackets, numbers, leading-whitespace, trailing-whitespace, double-space and special-placeholder)

		create
			--project-id string*
			--key-name string|json* 
			--description string
			--platforms array*
			--filenames json
			--tags array
			--comments array
			--screenshots json
			--translations json
			--is-plural bool
			--plural-name string
			--is-hidden bool
			--is-archived bool
			--context string
			--char-limit int
			--custom-attributes json

		retrieve
			--project-id string*
			--key-id int *
			--disable-references bool

		update
			--project-id string*
			--key-id int *
			--key-name string|json
			--description string
			--platforms json
			--filenames json
			--tags array
			--merge-tags bool
			--is-plural bool
			--plural-name string
			--is-hidden bool
			--is-archived bool
			--context string
			--char-limit int
			--custom-attributes json

		delete
			--project-id string*
			--key-id int *

	language
		list-system

		list
			--project-id string*

		create
			--project-id string*
			--lang-iso string*
			--custom-iso string
			--custom-name string
			--custom-plural-forms array

		retrieve
			--project-id string*
			--lang-id int*

		update
			--project-id string*
			--lang-id int*
			--lang-iso string
			--lang-name string
			--plural-forms array

		delete
			--project-id string*
			--lang-id int*

	order
		list
			--team-id int*
		
		create
			--team-id int*
			--project-id string*
			--card-id int*
			--briefing string*
			--source-language-iso string*
			--target-language-isos array*
			--keys array*
			--provider-slug string*
			--translation-tier int*
			--dry-run bool
			--translation-style string

		retrieve
			--team-id int*
			--order-id string

	payment-card
		list

		create
			--number int*
			--cvc int*
			--exp-month int*
			--exp-year int*

		retrieve
			--card-id

		delete
			--card-id

	project
		list
			--filter-team-id int
			--filter-names string
			--include-statistics bool
			--include-settings bool

		create
			--name string*
			--team-id int
			--description string
			--languages json
			--base-lang-iso string
			--project-type string

		retrieve
			--project-id string*

		update
			--project-id string*
			--name string
			--description string

		empty
			--project-id string*

		delete
			--project-id string*

	screenshot
		list
			--project-id string*

		create
			--project-id string*
			--file string*
			--title string
			--description string
			--ocr bool
			--key-ids array
			--tags array

		retrieve
			--project-id string*
			--screenshot-id int*

		update
			--project-id string*
			--screenshot-id int*
			--title string
			--description string
			--key-ids array
			--tags array

		delete
			--project-id string*
			--screenshot-id int*

	snapshot
		list
			--project-id string*

		create
			--project-id string*
			--title string

		restore
			--project-id string*
			--snapshot-id int*

		delete
			--project-id string*
			--snapshot-id int*

	task
		list
			--project-id string*
			--filter-title string

		create
			--project-id string*
			--description string
			--due-date string
			--keys array
			--languages json*
			--auto-close-languages bool
			--auto-close-task bool
			--initial-tm-leverage bool
			--task-type string
			--parent-task-id int
			--closing-tags array
			--do-lock-translations bool
			--custom-translation-status-ids array

		retrieve
			--project-id string*
			--task-id int

		update
			--project-id string*
			--task-id int
			--title string
			--description string
			--due-date string
			--languages json
			--auto-close-languages bool
			--auto-close-task bool
			--close-task bool
			--closing-tags array
			--do-lock-translations bool

		delete
			--project-id string*
			--task-id int

	team
		list

	team-user
		list
			--team-id int*

		retrieve
			--team-id int*
			--user-id int*

		update
			--team-id int*
			--user-id int*
			--role string

		delete
			--team-id int*
			--user-id int*

	team-user-group
		list
			--team-id int*

		create
			--team-id int*
			--name string*
			--is-reviewer bool*
			--is-admin bool*
			--admin-rights array
			--languages json

		retrieve
			--team-id int*
			--group-id int*

		update
			--team-id int*
			--group-id int*
			--name string*
			--is-reviewer string*
			--is-admin string*
			--admin-rights array
			--languages json

		add-projects
			--team-id int*
			--group-id int*
			--projects array

		remove-projects
			--team-id int*
			--group-id int*
			--projects array

		add-members
			--team-id int*
			--group-id int*
			--users array

		remove-members
			--team-id int*
			--group-id int*
			--users array

		delete
			--team-id int*
			--group-id int*

	translation
		list
			--project-id string*
			--disable-references bool
			--filter-lang-id string
			--filter-is-reviewed bool
			--filter-fuzzy bool
			--filter-qa-issues string

		retrieve
			--project-id string*
			--translation-id int*
			--disable-references bool

		update
			--project-id string*
			--translation-id int*
			--translation string*
			--is-fuzzy bool
			--is-reviewed bool			

	translation-provider
		list
			--team-id int*

		retrieve
			--team-id int*
			--provider-id int*

	translation-status
		list
			--project-id string*

		create
			--project-id string*
		 	--title string
		 	--color string

		retrieve
			--project-id string*
			--status-id int*

		update
			--project-id string*
			--status-id int*
			--title string
			--color string

		delete
			--project-id string*
			--status-id int*

		retrieve-colors



