[![Travis build status](https://travis-ci.org/Fitbit/smartling.svg?branch=master)](https://travis-ci.org/Fitbit/smartling)
[![Code Climate Maintainability](https://api.codeclimate.com/v1/badges/cccaa8f4dac2c3632a40/maintainability)](https://codeclimate.com/github/Fitbit/smartling)
[![Code Climate Coverage](https://api.codeclimate.com/v1/badges/cccaa8f4dac2c3632a40/test_coverage)](https://codeclimate.com/github/Fitbit/smartling)

# smartling

> CLI to work with Smartling translations

CLI tool provides simple unix-style command line interface to work with Smartling translations.

It is designed to make the downloading/uploading process very speedy.

## Features

- [x] Highly configurable via `.smartling.yml`
- [x] Designed to speedy upload/download (heavy) translations
- [x] Supports Smartling's API v2

## Commands

- [x] `push` - Uploads translations
- [x] `pull` - Downloads translations
- [x] `list` - Shows a list of local translations

## Homebrew

```
brew tap fitbit/tap
brew install smartling
```

## Download

All available releases you can find [here](https://github.com/Fitbit/smartling/releases).

## Usage

Please define `.smartling.yml` under your repo:

`.smartling.yml`

```yml
# The user identifier for your Smartling v2 API Token.
UserId: <SMARTLING_USER_ID>
# The token secret for your Smartling v2 API Token.
UserSecret: <SMARTLING_USER_SECRET>
# The project identifier for your Smartling v2 API Token.
ProjectId: <SMARTLING_PROJECT_ID>
# The alias is used in creating the "fileURI" that is registered with Smartling for uploaded files.
# A unique alias value is recommended but not required.
# The complete "fileURI" will be "ProjectAlias/FilePath".
ProjectAlias: <SMARTLING_PROJECT_ALIAS>
# List of files which will be uploaded/downloaded.
Files:
- # The Smartling API "fileType".
  # Possible values: "javaProperties", "ios", "android", "json" etc.
  # Please see documentation which types are supported:
  # https://help.smartling.com/docs/supported-file-types
  Type: json
  # "glob" expression defining which project files to upload to Smartling:
  # https://github.com/mattn/go-zglob
  # All files matching the expression will be uploaded.
  PathGlob: translations/**/en-US.json
  # List of "glob" expressions defining which project files will be excluded from upload.
  PathExclude:
  - translations/foo/en-US.json
  # The expression used to create a file path and name for translated files:
  # https://golang.org/pkg/html/template
  # Predefined variables for the expression are:
  # Path - original file path
  # Name - original file name without extension
  # Dir - original file directory path
  # Base - original file name with extension
  # Ext - original file extension
  # Locale - project locale code value
  # Also lot of built-in functions are available:
  # https://github.com/Masterminds/sprig
  PathExpression: '{{ .Dir }}/{{ .Locale }}{{ .Ext }}'
  # Defines whether uploaded content will automatically be authorized for translation.
  AuthorizeContent: true
  # File directives can be used to alter the way how Smartling handles your files.
  # Please see documentation which directives are available:
  # https://help.smartling.com/docs/supported-file-types
  Directives:
    # <directive name>: <directive value>
    string_format: NONE
# List of allowed locales and must list at least one locale.
# If you add extra locales to your project, you will need to update this file for the new locales.
# It also serves as a mapping of locale codes from Smartling API codes to the codes that are used in the project.
Locales:
  # <Smartling locale code>: <project locale code>
  de-DE: de-DE
  es-ES: es-ES
  fr-FR: fr-FR
  it-IT: it-IT
  ja-JP: ja-JP
  ko-KR: ko-KR
  pt-BR: pt-BR
  tr-TR: tr-TR
  zh-CN: zh-CN
  zh-TW: zh-TW
```

Then execute `smartling`:

`$ smartling`

```text
NAME:
   smartling - CLI to work with Smartling translations

USAGE:
   smartling [global options] command [command options] [arguments...]

COMMANDS:
     push      Uploads translations
     pull      Downloads translations
     list, ls  Shows a list of local translations
     help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --project-file value   Project configuration file (default: ".smartling.yml") [$SMARTLING_PROJECT_FILE]
   --project-id value     Project identifier for your Smartling v2 API Token [$SMARTLING_PROJECT_ID]
   --project-alias value  Unique alias of your project [$SMARTLING_PROJECT_ALIAS]
   --user-id value        User identifier for your Smartling v2 API Token [$SMARTLING_USER_ID]
   --user-secret value    Token secret for your Smartling v2 API Token [$SMARTLING_USER_SECRET]
   --no-color             Turn off colored output for log messages
   --verbose              Output verbose messages on internal operations
   --help, -h             Show help
   --version, -v          Output the version number
```

## Setup

- Run `brew install go glide`
- Run `make deps`

## Develop

- Run `make build` and execute `smartling`

## Test

- Run `make test`

## Cover

- Run `make cover` or `make cover-html`
