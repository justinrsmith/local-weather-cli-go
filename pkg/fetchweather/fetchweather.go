package fetchweather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Response struct
type Response struct {
	Weather  []Weather `json:"weather"`
	Main     Main      `json:"main"`
	Wind     []Wind    `json:"wind"`
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
	KelvinTemp    float64 `json:"temp"`
	Pressure      int     `json:"pressure"`
	Humidity      int     `json:"humidity"`
	KelvinTempMin float64 `json:"temp_min"`
	KelvinTempMax float64 `json:"temp_max"`
}

// Wind struct
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

// LocalWeather struct
type LocalWeather struct {
	City     string
	Current  string
	Temp     float64
	Humidity int
	High     float64
	Low      float64
}

func getData(url string) (int, []byte) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return response.StatusCode, responseData
}

// GetLocal function calls to the api and builds LocalWeather struct
func GetLocal(zip int) (LocalWeather, error) {
	_, data := getData(
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

	lw := LocalWeather{
		responseObject.Name,
		responseObject.Weather[0].Main,
		responseObject.Main.KelvinTemp,
		responseObject.Main.Humidity,
		responseObject.Main.KelvinTempMax,
		responseObject.Main.KelvinTempMin,
	}
	return lw, nil
}
