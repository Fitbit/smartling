[![Travis build status](http://img.shields.io/travis/Fitbit/smartling/master.svg?style=flat-square)](https://travis-ci.org/Fitbit/smartling)
[![Coverage Status](https://img.shields.io/coveralls/Fitbit/smartling/master.svg?style=flat-square)](https://coveralls.io/r/Fitbit/smartling?branch=master)

<a name="smartling"></a>
# smartling
> CLI to work with Smartling translations

CLI tool provides simple unix-style command line interface to work with Smartling translations.

It is designed to make the downloading/uploading process very speedy.

<a name="smartling-features"></a>
## Features

- [x] Highly configurable via `.smartling.yml`
- [x] Designed to speedy upload/download (heavy) translations
- [x] Supports Smartling's API v2

<a name="smartling-commands"></a>
## Commands

- [x] `push` - Uploads translations
- [x] `pull` - Downloads translations
- [x] `list` - Shows a list of local translations

<a name="smartling-homebrew"></a>
## Homebrew

```
brew tap mdreizin/tap
brew install smartling
```

<a name="smartling-download"></a>
## Download

All available releases you can find [here](https://github.com/Fitbit/smartling/releases).

<a name="smartling-usage"></a>
## Usage

Please define `.smartling.yml` under your repo:

`.smartling.yml`

```yml
# The User Identifier for your Smartling v2 API Token.
UserId: <SMARTLING_USER_ID>
# The token secret for your Smartling v2 API Token.
UserSecret: <SMARTLING_USER_SECRET>
# Unique id of your project.
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
  # http://docs.smartling.com/pages/supported-file-types
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
  # https://docs.smartling.com/pages/supported-file-types
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

```
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
   --project-file value   (default: ".smartling.yml") [$SMARTLING_PROJECT_FILE]
   --project-id value      [$SMARTLING_PROJECT_ID]
   --project-alias value   [$SMARTLING_PROJECT_ALIAS]
   --user-id value         [$SMARTLING_USER_ID]
   --user-secret value     [$SMARTLING_USER_SECRET]
   --no-color
   --help, -h             show help
   --version, -v          print the version

```

<a name="smartling-setup"></a>
## Setup

* Run `brew install go glide`
* Run `make restore`

<a name="smartling-develop"></a>
## Develop

* Run `make build` and execute `smartling`

<a name="smartling-test"></a>
## Test

* Run `make test`

<a name="smartling-cover"></a>
## Cover

* Run `make cover` or `make cover-html`
