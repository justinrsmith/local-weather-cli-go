# Local Weather CLI

Lightweight CLI for quickly checking the current weather in your terminal.

## Installation

### macOS

Local Weather CLI is available for installation via [Homebrew](https://brew.sh/):

```sh
brew install justinrsmith/local-weather-cli-go/local-weather
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

`-h` or `--help`

Get detailed information about CLI commands.

![Example](./docs/examples/help_usage.svg)
