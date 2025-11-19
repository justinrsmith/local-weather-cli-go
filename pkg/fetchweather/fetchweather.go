package fetchweather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
)

// Response represents the API response structure.
type Response struct {
	Weather  []Weather `json:"weather"`
	Main     Main      `json:"main"`
	Wind     Wind      `json:"wind"`
	Name     string    `json:"name"`
	RespCode int       `json:"cod"`
	Message  string    `json:"message"`
}

// Weather struct represents weather condition information.
type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
}

// Main represents the main weather data including temperature, pressure, and humidity.
type Main struct {
	KelvinTemp      float64 `json:"temp"`
	KelvinFeelsLike float64 `json:"feels_like"`
	Pressure        int     `json:"pressure"`
	Humidity        int     `json:"humidity"`
}

// Wind represents wind speed and direction data.
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

// LocalWeather represents the processed weather data for display.
type LocalWeather struct {
	City          string
	Current       string
	Temp          float64
	Humidity      int
	WindSpeed     float64
	WindDirection string
	FeelsLike     float64
}

var zgetData = getData

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

// getData fetches data from the given URL and returns the response body.
func getData(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected HTTP status: %d", response.StatusCode)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return responseData, nil
}

// WindDegreeToDirection converts a wind degree (0-360) to a cardinal direction string.
// Returns one of 16 cardinal and intercardinal directions.
func WindDegreeToDirection(degree float64) string {
	directions := [16]string{
		"North", "North-Northeast", "Northeast", "East-Northeast", "East",
		"East-Southeast", "Southeast", "South-Southeast", "South",
		"South-Southwest", "Southwest", "West-Southwest", "West",
		"West-Northwest", "Northwest", "North-Northwest"}

	index := int((degree+11.25)/22.5) % 16

	return directions[index]
}

// GetLocal fetches weather data for the given zip code and returns a LocalWeather struct.
// The scale parameter must be one of: "C" (Celsius), "F" (Fahrenheit), or "K" (Kelvin).
func GetLocal(zip int, scale string) (LocalWeather, error) {
	data, err := zgetData(
		fmt.Sprintf(
			"https://local-weather-api-456015.appspot.com/?zip=%d,us",
			zip,
		),
	)
	if err != nil {
		return LocalWeather{}, fmt.Errorf("failed to get weather data: %w", err)
	}

	var responseObject Response
	if err := json.Unmarshal(data, &responseObject); err != nil {
		return LocalWeather{}, fmt.Errorf("failed to parse response: %w", err)
	}

	if responseObject.RespCode != 200 {
		return LocalWeather{}, errors.New(responseObject.Message)
	}

	if len(responseObject.Weather) == 0 {
		return LocalWeather{}, errors.New("no weather data in response")
	}

	var currentTemp float64
	var feelsLikeTemp float64

	switch {
	case scale == "K":
		currentTemp = responseObject.Main.KelvinTemp
		feelsLikeTemp = responseObject.Main.KelvinFeelsLike
	case scale == "C":
		currentTemp = responseObject.Main.KelvinTemp - 273.15
		feelsLikeTemp = responseObject.Main.KelvinFeelsLike - 273.15
	case scale == "F":
		currentTemp = responseObject.Main.KelvinTemp*9/5 - 459.67
		feelsLikeTemp = responseObject.Main.KelvinFeelsLike*9/5 - 459.67
	}

	lw := LocalWeather{
		responseObject.Name,
		responseObject.Weather[0].Main,
		toFixed(currentTemp, 2),
		responseObject.Main.Humidity,
		responseObject.Wind.Speed,
		WindDegreeToDirection(float64(responseObject.Wind.Deg)),
		toFixed(feelsLikeTemp, 2),
	}
	return lw, nil
}
