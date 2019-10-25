package main

import (
	"os"

	cmd "github.com/justinrsmith/local-weather-cli-go/cmd/local-weather"
)

func main() {
	cmd.Run(os.Args)
}
