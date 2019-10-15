# Lokalise CLI v2

## Getting started
Lokalise CLI v2 allows you to manipulate any object or data in your Lokalise workspace, which includes files, projects, keys, translations, comments, contributors, teams and more. All endpoints available in [Lokalise API v2](https://lokalise.com/api2docs/curl/) are accessible using this tool.

## Installation
Get the binaries for your platform, unarchive and put into any executable folder. All set! 

### MacOS
Install using Homebrew:
```
brew tap lokalise/cli-2
brew install lokalise2
```
or get the binaries:
- [lokalise2-2.00-darwin-amd64.tgz](https://s3-eu-west-1.amazonaws.com/lokalise-assets/cli2/lokalise2-2.00-darwin-amd64.tgz)
- [lokalise2-2.00-darwin-386.tgz](https://s3-eu-west-1.amazonaws.com/lokalise-assets/cli2/lokalise2-2.00-darwin-386.tgz)

### Linux
- [lokalise2-2.00-linux-amd64.tgz](https://s3-eu-west-1.amazonaws.com/lokalise-assets/cli2/lokalise2-2.00-linux-amd64.tgz)
- [lokalise2-2.00-linux-386.tgz](https://s3-eu-west-1.amazonaws.com/lokalise-assets/cli2/lokalise2-2.00-linux-386.tgz)

### FreeBSD
- [lokalise2-2.00-freebsd-amd64.tgz](https://s3-eu-west-1.amazonaws.com/lokalise-assets/cli2/lokalise2-2.00-freebsd-amd64.tgz)
- [lokalise2-2.00-freebsd-386.tgz](https://s3-eu-west-1.amazonaws.com/lokalise-assets/cli2/lokalise2-2.00-freebsd-386.tgz)

### Windows
- [lokalise2-2.00-windows-amd64.tgz](https://s3-eu-west-1.amazonaws.com/lokalise-assets/cli2/lokalise2-2.00-windows-amd64.tgz)
- [lokalise2-2.00-windows-386.tgz](https://s3-eu-west-1.amazonaws.com/lokalise-assets/cli2/lokalise2-2.00-windows-386.tgz)

### Docker
See [DockerHub](https://hub.docker.com/r/lokalise/lokalise-cli-2) for more information.

## Tokens
All endpoints require the `--token` parameter. You can generate your API token in [Personal profile](https://lokalise.com/profile). Note, the token is personal and mimics your access level on team and project level. If you require a user-independent API token, create a separate user in your team (e.g. API user), set proper access rights and use this user's token.

## Passing the flags

#### Boolean
`boolean` flags *must* be passed via the `=` sign, i.e. `--original-filenames=false`. 

#### Strings
Parameters of type `strings` should be comma-delimited, i.e. `--include-tags=one,two`. 

#### JSON objects
Some parameters require a JSON-encoded object passed as string, i.e. `--languages='"[{\"lang_iso\":\"en\",\"custom_iso\":\"en-us\"},{\"lang_iso\":\"en_GB\",\"custom_iso\":\"en-gb\"}]"'`.

## Usage

Run `lokalise2 --help` from the terminal to see the list of commands and subcommands. Run `lokalise2 command subcommand --help` to see the help page.

### Commands

* [lokalise2 branch](docs/lokalise2_branch.md)	 - Manage branches
* [lokalise2 comment](docs/lokalise2_comment.md)	 - Manage key comments
* [lokalise2 contributor](docs/lokalise2_contributor.md)	 - Manage project contributors
* [lokalise2 file](docs/lokalise2_file.md)	 - Upload and download files
* [lokalise2 key](docs/lokalise2_key.md)	 - Manage keys
* [lokalise2 language](docs/lokalise2_language.md)	 - Manage languages
* [lokalise2 order](docs/lokalise2_order.md)	 - Manage orders
* [lokalise2 payment-card](docs/lokalise2_payment-card.md)	 - Manage payment cards
* [lokalise2 project](docs/lokalise2_project.md)	 - Manage projects
* [lokalise2 screenshot](docs/lokalise2_screenshot.md)	 - Manage screenshots
* [lokalise2 snapshot](docs/lokalise2_snapshot.md)	 - Manage snapshots
* [lokalise2 task](docs/lokalise2_task.md)	 - Manage tasks
* [lokalise2 team](docs/lokalise2_team.md)	 - List teams
* [lokalise2 team-user](docs/lokalise2_team-user.md)	 - Manage team users
* [lokalise2 team-user-group](docs/lokalise2_team-user-group.md)	 - Manage team user groups
* [lokalise2 translation](docs/lokalise2_translation.md)	 - Manage translations
* [lokalise2 translation-provider](docs/lokalise2_translation-provider.md)	 - List translation providers
* [lokalise2 translation-status](docs/lokalise2_translation-status.md)	 - Manage custom translation statuses
* [lokalise2 webhook](docs/lokalise2_webhook.md)	 - Manage webhooks
