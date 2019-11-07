package fetchweather

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetData(t *testing.T) {
	mockResp := `{"hello": "world"}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, mockResp)
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	resp, err := getData(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	if strings.TrimSuffix(string(resp), "\n") != mockResp {
		fmt.Println(data)
		t.Errorf("Incorrect response returned")
	}
}

func TestGetLocalKelvin(t *testing.T) {
	old := zgetData
	defer func() { zgetData = old }()
	zgetData = func(url string) ([]byte, error) {
		mockResp := `{"coord":{"lon":-89.5,"lat":43.03},"weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04d"}],"base":"stations","main":{"temp":272.84,"pressure":1014,"humidity":82,"temp_min":271.48,"temp_max":274.26},"visibility":16093,"wind":{"speed":7.7,"deg":310,"gust":11.3},"clouds":{"all":90},"dt":1572552019,"sys":{"type":1,"id":3551,"country":"US","sunrise":1572525057,"sunset":1572562328},"timezone":-18000,"id":0,"name":"Madison","cod":200}`
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintln(w, mockResp)
		}))
		defer ts.Close()

		res, err := http.Get(ts.URL)
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		return data, err
	}

	localWeather, _ := GetLocal(61008, "K")
	if localWeather.Temp != 272.84 {
		t.Errorf("Temperature is not the correct Kelvin temperature")
	}
}

func TestGetLocalCelsius(t *testing.T) {
	old := zgetData
	defer func() { zgetData = old }()
	zgetData = func(url string) ([]byte, error) {
		mockResp := `{"coord":{"lon":-89.5,"lat":43.03},"weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04d"}],"base":"stations","main":{"temp":272.84,"pressure":1014,"humidity":82,"temp_min":271.48,"temp_max":274.26},"visibility":16093,"wind":{"speed":7.7,"deg":310,"gust":11.3},"clouds":{"all":90},"dt":1572552019,"sys":{"type":1,"id":3551,"country":"US","sunrise":1572525057,"sunset":1572562328},"timezone":-18000,"id":0,"name":"Madison","cod":200}`
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintln(w, mockResp)
		}))
		defer ts.Close()

		res, err := http.Get(ts.URL)
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		return data, err
	}

	localWeather, _ := GetLocal(61008, "C")
	if localWeather.Temp != -0.31 {
		t.Errorf("Temperature is not the correct Celsius temperature")
	}
}

func TestGetLocalFahrenheit(t *testing.T) {
	old := zgetData
	defer func() { zgetData = old }()
	zgetData = func(url string) ([]byte, error) {
		mockResp := `{"coord":{"lon":-89.5,"lat":43.03},"weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04d"}],"base":"stations","main":{"temp":272.84,"pressure":1014,"humidity":82,"temp_min":271.48,"temp_max":274.26},"visibility":16093,"wind":{"speed":7.7,"deg":310,"gust":11.3},"clouds":{"all":90},"dt":1572552019,"sys":{"type":1,"id":3551,"country":"US","sunrise":1572525057,"sunset":1572562328},"timezone":-18000,"id":0,"name":"Madison","cod":200}`
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintln(w, mockResp)
		}))
		defer ts.Close()

		res, err := http.Get(ts.URL)
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		return data, err
	}

	localWeather, _ := GetLocal(61008, "F")
	if localWeather.Temp != 31.44 {
		t.Errorf("Temperature is not the correct Fahrenheit temperature")
	}
}

func TestGetLocalFails(t *testing.T) {
	old := zgetData
	defer func() { zgetData = old }()
	zgetData = func(url string) ([]byte, error) {
		mockResp := `{"cod":404,"message":"city not found"}`
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintln(w, mockResp)
		}))
		defer ts.Close()

		res, err := http.Get(ts.URL)
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		return data, err
	}

	_, err := GetLocal(6100, "K")
	if err == nil {
		t.Errorf("Error was returned for non existing city")
	}
}

func TestRound(t *testing.T) {
	expected := 2
	got := round(1.674)

	if got != expected {
		t.Errorf("Rounding failed")
	}
}

func TestToFixed(t *testing.T) {
	expected := 1.65
	got := toFixed(1.647, 2)

	if got != expected {
		t.Errorf("Rounding to fixed decimal failed")
	}
}
