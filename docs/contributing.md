# Contributing to Local Weather CLI

All contributions to the Local Weather CLI are welcome and appreciated!

## Development Setup

The Local Weather CLI is written in Go. To download the source code:

```sh
go get -u github.com/justinrsmith/local-weather-cli-go/...
```

After downloading, `cd` to the directory:

```sh
cd go/src/github.com/justinrsmith/local-weather-cli-go
```

Before starting any development, run the test suite to make sure everything is passing as expected:

```sh
make test
```

If the tests passed you can build the application with:

```sh
make build
```

Then you can run the local version of the CLI:

```sh
local-weather --zipcode 10007
```

## Tests

It is expected that any submitted pull requests will include proper unit testing.
