# Local Weather CLI

Lightweight CLI for quickly checking the current weather

## Installation

To install latest binary release, run:

```
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
