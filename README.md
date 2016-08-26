[![Travis build status](http://img.shields.io/travis/mdreizin/smartling/master.svg?style=flat-square)](https://travis-ci.org/mdreizin/smartling)
[![Coverage Status](https://img.shields.io/coveralls/mdreizin/smartling/master.svg?style=flat-square)](https://coveralls.io/r/mdreizin/smartling?branch=master)

<h1 id="smartling">smartling</h1>
> Smartling CLI to `upload` and `download` translations

Supports only Smartling's API v2

<h2 id="smartling-commands">Commands</h2>

- [x] `push` - Uploads translations
- [x] `pull` - Downloads translations
- [x] `list` - Displays list of local translations

<h2 id="smartling-download">Download</h2>

All available releases you can find [here](https://github.com/mdreizin/smartling/releases).

<h2 id="smartling-usage">Usage</h2>

Please defined `.smartling.yml` under your repo:

`.smartling.yml`

```yml
UserId: <SMARTLING_USER_ID>
UserSecret: <SMARTLING_USER_SECRET>
ProjectId: <SMARTLING_PROJECT_ID>
ProjectAlias: <SMARTLING_PROJECT_ALIAS>
Files:
- Type: json
  PathGlob: translations/**/en-US.json
  PathExclude:
  - translations/foo/en-US.json
  PathExpression: '{{ .Dir }}/{{ .Locale }}{{ .Ext }}'
  AuthorizeContent: true
  Directives:
    string_format: NONE
Locales:
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
   smartling - Smartling CLI to `upload` and `download` translations

USAGE:
   smartling [global options] command [command options] [arguments...]

VERSION:
   dev

AUTHOR(S):
   Marat Dreizin <marat.dreizin@gmail.com>

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
   --help, -h             show help
   --version, -v          print the version

```

<h2 id="smartling-setup">Setup</h2>

* Run `brew install go glide`
* Run `make restore`

<h2 id="smartling-develop">Develop</h2>

* Run `make build` and execute `smartling`

<h2 id="smartling-test">Test</h2>

* Run `make test`

<h2 id="smartling-cover">Cover</h2>

* Run `make cover` or `make cover-html`
