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

func getTempScaleLabel(scale string) string {
	switch {
	case scale == "K":
		return "K"
	case scale == "F":
		return "\u00b0F"
	case scale == "C":
		return "\u00b0C"
	}
	return ""
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

	tempScale := getTempScaleLabel(scale)

	localWeatherStr := []string{
		localWeather.City,
		localWeather.Current,
		fmt.Sprintf("%s%s", floatToString(localWeather.Temp), tempScale),
		fmt.Sprintf("%s%%", strconv.Itoa(localWeather.Humidity)),
		fmt.Sprintf("%s%s", floatToString(localWeather.High), tempScale),
		fmt.Sprintf("%s%s", floatToString(localWeather.Low), tempScale),
	}

	generateOutput(os.Stdout, localWeatherStr)

	return nil
}

func main() {
	if err := execute(); err != nil {
		os.Exit(-1)
	}
}
