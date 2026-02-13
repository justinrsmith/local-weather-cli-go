package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/justinrsmith/local-weather-cli-go/v3/pkg/fetchweather"
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
		"Wind",
		"Feels Like",
	})

	table.Append(data)
	table.Render()
}

func getTempScaleLabel(scale string) string {
	switch scale {
	case "K":
		return "K"
	case "F":
		return "\u00b0F"
	case "C":
		return "\u00b0C"
	}
	return ""
}

func execute() error {
	var zipcode int
	var scale string

	flag.Usage = func() {
		flag.PrintDefaults()
	}

	flag.IntVar(&zipcode, "zipcode", 0, "zip code of U.S. city")
	flag.IntVar(&zipcode, "z", 0, "zip code of U.S. city shorthand")
	flag.StringVar(&scale, "scale", "F", "temperature scale to use {C - Celsius|F - Fahrenheit)|K - Kelvin}")
	flag.StringVar(&scale, "s", "F", "temperature scale to use {C - Celsius|F - Fahrenheit)|K - Kelvin}")

	flag.Parse()
	if zipcode == 0 {
		flag.PrintDefaults()
		return fmt.Errorf("zipcode is required")
	}
	// temp scale flag should be case insensitive
	scale = strings.ToUpper(scale)
	if !slices.Contains([]string{"C", "K", "F"}, scale) {
		flag.PrintDefaults()
		return fmt.Errorf("invalid temperature scale: %s", scale)
	}

	localWeather, err := fetchweather.GetLocal(zipcode, scale)
	if err != nil {
		return err
	}

	tempScale := getTempScaleLabel(scale)

	localWeatherStr := []string{
		localWeather.City,
		localWeather.Current,
		fmt.Sprintf("%s%s", floatToString(localWeather.Temp), tempScale),
		fmt.Sprintf("%s%%", strconv.Itoa(localWeather.Humidity)),
		fmt.Sprintf("%smph %s", floatToString(localWeather.WindSpeed), localWeather.WindDirection),
		fmt.Sprintf("%s%s", floatToString(localWeather.FeelsLike), tempScale),
	}

	generateOutput(os.Stdout, localWeatherStr)

	return nil
}

func main() {
	if err := execute(); err != nil {
		log.Fatal(err)
	}
}
