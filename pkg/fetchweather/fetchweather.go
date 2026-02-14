package fetchweather

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"time"
)

// Response represents the raw API response from the weather service.
type Response struct {
	Weather  []Weather `json:"weather"`
	Main     Main      `json:"main"`
	Wind     Wind      `json:"wind"`
	Name     string    `json:"name"`
	RespCode int       `json:"cod"`
	Message  string    `json:"message"`
}

// Weather represents a weather condition entry from the API.
type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
}

// Main represents the primary weather measurements.
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

// LocalWeather holds the processed weather data for display.
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

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return math.Round(num*output) / output
}

var httpClient = &http.Client{Timeout: 10 * time.Second}

func getData(url string) ([]byte, error) {
	response, err := httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetching weather data: %w", err)
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}

	return responseData, nil
}

func WindDegreeToDirection(degree float64) string {
	directions := [16]string{
		"North", "North-Northeast", "Northeast", "East-Northeast", "East",
		"East-Southeast", "Southeast", "South-Southeast", "South",
		"South-Southwest", "Southwest", "West-Southwest", "West",
		"West-Northwest", "Northwest", "North-Northwest"}

	index := int((degree+11.25)/22.5) % 16

	return directions[index]
}

// GetLocal calls the weather API and returns processed weather data.
func GetLocal(zip int, scale string) (LocalWeather, error) {
	data, err := zgetData(
		fmt.Sprintf(
			"https://local-weather-api-456015.appspot.com/?zip=%d,us",
			zip,
		),
	)
	if err != nil {
		return LocalWeather{}, err
	}

	var responseObject Response
	if err := json.Unmarshal(data, &responseObject); err != nil {
		return LocalWeather{}, fmt.Errorf("parsing weather response: %w", err)
	}

	if responseObject.RespCode != 200 {
		return LocalWeather{}, fmt.Errorf("weather API error: %s", responseObject.Message)
	}

	var currentTemp float64
	var feelsLikeTemp float64

	switch scale {
	case "K":
		currentTemp = responseObject.Main.KelvinTemp
		feelsLikeTemp = responseObject.Main.KelvinFeelsLike
	case "C":
		currentTemp = responseObject.Main.KelvinTemp - 273.15
		feelsLikeTemp = responseObject.Main.KelvinFeelsLike - 273.15
	case "F":
		currentTemp = responseObject.Main.KelvinTemp*9/5 - 459.67
		feelsLikeTemp = responseObject.Main.KelvinFeelsLike*9/5 - 459.67
	}

	lw := LocalWeather{
		City:          responseObject.Name,
		Current:       responseObject.Weather[0].Main,
		Temp:          toFixed(currentTemp, 2),
		Humidity:      responseObject.Main.Humidity,
		WindSpeed:     responseObject.Wind.Speed,
		WindDirection: WindDegreeToDirection(float64(responseObject.Wind.Deg)),
		FeelsLike:     toFixed(feelsLikeTemp, 2),
	}
	return lw, nil
}
