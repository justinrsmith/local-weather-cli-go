![](https://github.com/justinrsmith/local-weather-cli-go/workflows/build/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/justinrsmith/local-weather-cli-go)](https://goreportcard.com/report/github.com/justinrsmith/local-weather-cli-go)

# Local Weather CLI

Lightweight CLI for quickly checking the current weather in your terminal.

## Installation

### macOS

Local Weather CLI is available for installation on macOS via [Homebrew](https://brew.sh/):

```sh
brew install justinrsmith/local-weather-cli-go/local-weather
```

### Windows

Local Weather CLI is available for installation on Windows via [Scoop](https://scoop.sh/):

```sh
scoop bucket add local-weather https://github.com/justinrsmith/local-weather-cli-go.git
scoop install local-weather
```

### Docker

Local Weather CLI is also available as a [Docker](https://www.docker.com/) image: [`justinrsmith88/local-weather-cli-go`](https://hub.docker.com/r/justinrsmith88/local-weather-cli-go):

```sh
docker run --rm justinrsmith88/local-weather-cli-go:latest --help
```

### Binary

To install the latest binary release, run:

```sh
curl --location --silent "https://github.com/justinrsmith/local-weather-cli-go/releases/download/v<VERSION>/local-weather-cli-go_<VERSION>_Darwin_x86_64.tar.gz" | tar xz -C /tmp
sudo mv /tmp/local-weather-cli-go /usr/local/bin
```

## Usage

Upon installation of the CLI you can quickly check a city's weather.

```sh
local-weather --zipcode 10007
```

### Options

`-z <zipcode>` or `--zipcode <zipcode>`

Get the weather for a US city by zipcode.

![Example](./docs/examples/zipcode_usage.svg)

`-s <scale to use>` or `--scale <scale to use>`

Select which temperature scale to use in output.

Valid scales:
- `C` - Celcius
- `F` - Fahrenheit (default)
- `K` - Kelvin

![Example](./docs/examples/scale_usage.svg)

`-h` or `--help`

Get detailed information about CLI commands.

![Example](./docs/examples/help_usage.svg)
