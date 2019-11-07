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

	flag.IntVar(&zipcode, "zipcode", 0, "zip code of city")
	flag.IntVar(&zipcode, "z", 0, "zip code of city shorthand")

	flag.Parse()
	if zipcode == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	localWeather, err := fetchweather.GetLocal(zipcode, "F")
	if err != nil {
		log.Fatal(err)
	}

	localWeatherStr := []string{
		localWeather.City,
		localWeather.Current,
		floatToString(localWeather.Temp),
		strconv.Itoa(localWeather.Humidity),
		floatToString(localWeather.High),
		floatToString(localWeather.Low),
	}

	generateOutput(os.Stdout, localWeatherStr)

	return nil
}

func main() {
	if err := execute(); err != nil {
		os.Exit(-1)
	}
}
