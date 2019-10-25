package fetchweather

import "testing"

func TestGetData(t *testing.T) {
	status, _ := getData("https://local-weather-api-256018.appspot.com/?zip=61008,us")
	if status != 200 {
		t.Errorf("Bad response from server")
	}
}

func TestErrorRaised(t *testing.T) {
	_, err := GetLocal("6100")
	if err == nil {
		t.Errorf("Error was not returned for non existing city")
	}
}

func TestLocal(t *testing.T) {
	localWeather, _ := GetLocal("61008")
	if localWeather.City != "Belvidere" {
		t.Errorf("Incorrect data")
	}
}
