package fetchweather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
)

// Response struct
type Response struct {
	Weather  []Weather `json:"weather"`
	Main     Main      `json:"main"`
	Wind     Wind      `json:"wind"`
	Name     string    `json:"name"`
	RespCode int       `json:"cod"`
	Message  string    `json:"message"`
}

// Weather struct
type Weather struct {
	ID          string `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
}

// Main struct
type Main struct {
	KelvinTemp      float64 `json:"temp"`
	KelvinFeelsLike float64 `json:"feels_like"`
	Pressure        int     `json:"pressure"`
	Humidity        int     `json:"humidity"`
}

// Wind struct
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

// LocalWeather struct
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

func getData(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return responseData, err
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

// GetLocal function calls to the api and builds LocalWeather struct
func GetLocal(zip int, scale string) (LocalWeather, error) {
	data, _ := zgetData(
		fmt.Sprintf(
			"https://local-weather-api-256018.appspot.com/?zip=%d,us",
			zip,
		),
	)
	var responseObject Response
	json.Unmarshal(data, &responseObject)

	if responseObject.RespCode != 200 {
		return LocalWeather{}, errors.New(responseObject.Message)
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
