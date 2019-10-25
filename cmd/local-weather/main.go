package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/justinrsmith/local-weather-cli-go/pkg/fetchweather"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

func floatToString(inputNum float64) string {
	return fmt.Sprintf("%.1f", inputNum)
}

func kelvinToFarhenheit(temp float64) float64 {
	farhenheit := temp*9/5 - 459.67
	return farhenheit
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

func Run(args []string) {
	var zipcode string

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "zipcode, z",
			Usage:       "zipcode",
			Destination: &zipcode,
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.NumFlags() < 1 {
			cli.ShowAppHelp(c)
			return cli.NewExitError("You must provide a zipcode", 2)
		}

		localWeather, err := fetchweather.GetLocal(zipcode)
		if err != nil {
			return cli.NewExitError(err, 2)
		}

		localWeatherStr := []string{
			localWeather.City,
			localWeather.Current,
			floatToString(kelvinToFarhenheit(localWeather.Temp)),
			strconv.Itoa(localWeather.Humidity),
			floatToString(kelvinToFarhenheit(localWeather.High)),
			floatToString(kelvinToFarhenheit(localWeather.Low)),
		}
		generateOutput(os.Stdout, localWeatherStr)

		return nil
	}

	err := app.Run(args)
	if err != nil {
		log.Fatal(err)
	}
}

// func main() {
// 	run(os.Args)
// }
