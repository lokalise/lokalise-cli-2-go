# Lokalise CLI v2

## Getting started

Lokalise CLI v2 allows you to manipulate any object or data in your Lokalise workspace, which includes files, projects, keys, translations, comments, contributors, teams and more. All endpoints available in [Lokalise API v2](https://lokalise.com/api2docs/curl/) are accessible using this tool.

## Installation

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

## Usage

Refer to the [CLI tool reference](docs/lokalise2.md) for usage.