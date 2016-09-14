{"gitdown": "badge", "name": "travis"}
{"gitdown": "badge", "name": "coveralls"}

# {"gitdown": "gitinfo", "name": "name"}
> Smartling CLI to `upload` and `download` translations

Supports only Smartling's API v2

## Commands

- [x] `push` - Uploads translations
- [x] `pull` - Downloads translations
- [x] `list` - Shows a list of local translations

## Homebrew

```
brew tap mdreizin/tap
brew install smartling
```

## Download

All available releases you can find [here]({"gitdown": "gitinfo", "name": "url"}/releases).

## Usage

Please defined `.smartling.yml` under your repo:

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
* Run `make restore`

## Develop

* Run `make build` and execute `smartling`

## Test

* Run `make test`

## Cover

* Run `make cover` or `make cover-html`
