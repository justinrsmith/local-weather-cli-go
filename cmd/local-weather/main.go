package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/justinrsmith/local-weather-cli-go/pkg/fetchweather"
	"github.com/olekukonko/tablewriter"
)

func floatToString(inputNum float64) string {
	return fmt.Sprintf("%.1f", inputNum)
}

func generateOutput(dst io.Writer, data []string) {
	table := tablewriter.NewWriter(dst)
	table.SetHeader([]string{
		"City",
		"Current Condition",
		"Current Temp",
		"Humidity",
		"High Temp",
		"Low Temp",
	})

	table.Append(data)
	table.Render() // Send output
}

func execute() error {
	var zipcode int
	var scale string

	flag.Usage = func() {
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.IntVar(&zipcode, "zipcode", 0, "zip code of U.S. city")
	flag.IntVar(&zipcode, "z", 0, "zip code of U.S. city shorthand")
	flag.StringVar(&scale, "scale", "F", "temperature scale to use {C - Celcius|F - Fahrenheit)|K - Kelvin}")
	flag.StringVar(&scale, "s", "F", "temperature scale to use {C - Celcius|F - Fahrenheit)|K - Kelvin}")

	flag.Parse()
	if zipcode == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	localWeather, err := fetchweather.GetLocal(zipcode, scale)
	if err != nil {
		log.Fatal(err)
	}

	localWeatherStr := []string{
		localWeather.City,
		localWeather.Current,
		fmt.Sprintf("%s\u00b0F", floatToString(localWeather.Temp)),
		fmt.Sprintf("%s%%", strconv.Itoa(localWeather.Humidity)),
		fmt.Sprintf("%s\u00b0F", floatToString(localWeather.High)),
		fmt.Sprintf("%s\u00b0F", floatToString(localWeather.Low)),
	}

	generateOutput(os.Stdout, localWeatherStr)

	return nil
}

func main() {
	if err := execute(); err != nil {
		os.Exit(-1)
	}
}
