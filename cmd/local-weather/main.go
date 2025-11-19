package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/justinrsmith/local-weather-cli-go/v3/pkg/fetchweather"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	zipcode int
	scale   string
	verbose bool
)

// floatToString formats a float64 to a string with one decimal place.
func floatToString(inputNum float64) string {
	return fmt.Sprintf("%.1f", inputNum)
}

// stringInSlice checks if a string exists in a slice of strings.
func stringInSlice(str string, list []string) bool {
	for _, val := range list {
		if str == val {
			return true
		}
	}
	return false
}

// generateOutput creates a formatted table output with weather data.
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

// getTempScaleLabel returns the appropriate label for the temperature scale.
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

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "local-weather",
	Short: "Lightweight CLI for quickly checking the current weather in your terminal",
	Long: `Local Weather CLI is a lightweight command-line tool for quickly checking
the current weather in your terminal. Currently supports U.S. zipcodes.`,
	RunE: runWeather,
}

// runWeather executes the weather fetching and display logic.
func runWeather(cmd *cobra.Command, args []string) error {
	// Setup structured logging
	var logger *slog.Logger
	if verbose {
		logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelError, // Only log errors by default
		}))
	}

	// Validate zipcode
	if zipcode == 0 {
		return fmt.Errorf("zipcode is required")
	}

	// Validate and normalize temperature scale
	scale = strings.ToUpper(scale)
	validScales := []string{"C", "K", "F"}
	if !stringInSlice(scale, validScales) {
		return fmt.Errorf("invalid temperature scale: %s (valid options: C, F, K)", scale)
	}

	// Fetch weather data
	ctx := cmd.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	// Add logger to context
	ctx = context.WithValue(ctx, fetchweather.LoggerKey{}, logger)

	logger.Debug("Fetching weather data", "zipcode", zipcode, "scale", scale)
	localWeather, err := fetchweather.GetLocal(ctx, zipcode, scale)
	if err != nil {
		logger.Error("Failed to fetch weather", "error", err, "zipcode", zipcode)
		return fmt.Errorf("failed to fetch weather: %w", err)
	}
	logger.Debug("Successfully fetched weather data", "city", localWeather.City)

	// Format output
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

func init() {
	rootCmd.Flags().IntVarP(&zipcode, "zipcode", "z", 0, "zip code of U.S. city")
	rootCmd.Flags().StringVarP(&scale, "scale", "s", "F", "temperature scale to use (C - Celsius, F - Fahrenheit, K - Kelvin)")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose logging")

	// Mark zipcode as required
	rootCmd.MarkFlagRequired("zipcode")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
