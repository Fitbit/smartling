{"gitdown": "badge", "name": "travis"}
{"gitdown": "badge", "name": "coveralls"}

# {"gitdown": "gitinfo", "name": "name"}
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

All available releases you can find [here]({"gitdown": "gitinfo", "name": "url"}/releases).

## Usage

Please define `.smartling.yml` under your repo:

`.smartling.yml`

```yml
{"gitdown": "include", "file": "smartling.yml"}
```

Then execute `smartling`:

`$ smartling`

```
{"gitdown": "include", "file": "cli.txt"}
```

## Setup

* Run `brew install go glide`
* Run `make deps`

## Develop

* Run `make build` and execute `smartling`

## Test

* Run `make test`

## Cover

* Run `make cover` or `make cover-html`
