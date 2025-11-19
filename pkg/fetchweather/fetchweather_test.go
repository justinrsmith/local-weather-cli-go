package fetchweather

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// mockFetcher is a test implementation of DataFetcher.
type mockFetcher struct {
	fetchFunc func(ctx context.Context, url string) ([]byte, error)
}

// Fetch implements DataFetcher interface.
func (m *mockFetcher) Fetch(ctx context.Context, url string) ([]byte, error) {
	return m.fetchFunc(ctx, url)
}

func TestGetData(t *testing.T) {
	mockResp := `{"hello": "world"}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, mockResp)
	}))
	defer ts.Close()

	ctx := context.Background()
	resp, err := getData(ctx, ts.URL)
	if err != nil {
		t.Fatalf("getData failed: %v", err)
	}

	if string(resp) != mockResp {
		t.Errorf("Incorrect response returned: got %q, want %q", string(resp), mockResp)
	}
}

func TestGetLocalKelvin(t *testing.T) {
	mockResp := `{"coord":{"lon":-89.5,"lat":43.03},"weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04d"}],"base":"stations","main":{"temp":272.84,"pressure":1014,"humidity":82,"temp_min":271.48,"temp_max":274.26},"visibility":16093,"wind":{"speed":7.7,"deg":310,"gust":11.3},"clouds":{"all":90},"dt":1572552019,"sys":{"type":1,"id":3551,"country":"US","sunrise":1572525057,"sunset":1572562328},"timezone":-18000,"id":0,"name":"Madison","cod":200}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, mockResp)
	}))
	defer ts.Close()

	fetcher := &mockFetcher{
		fetchFunc: func(ctx context.Context, url string) ([]byte, error) {
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, ts.URL, nil)
			if err != nil {
				return nil, err
			}
			client := &http.Client{}
			res, err := client.Do(req)
			if err != nil {
				return nil, err
			}
			defer res.Body.Close()
			data, err := io.ReadAll(res.Body)
			if err != nil {
				return nil, err
			}
			return data, nil
		},
	}

	ctx := context.Background()
	localWeather, err := GetLocal(ctx, fetcher, 61008, "K")
	if err != nil {
		t.Fatalf("GetLocal failed: %v", err)
	}
	if localWeather.Temp != 272.84 {
		t.Errorf("Temperature is not the correct Kelvin temperature: got %f, want 272.84", localWeather.Temp)
	}
}

func TestGetLocalCelsius(t *testing.T) {
	mockResp := `{"coord":{"lon":-89.5,"lat":43.03},"weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04d"}],"base":"stations","main":{"temp":272.84,"pressure":1014,"humidity":82,"temp_min":271.48,"temp_max":274.26},"visibility":16093,"wind":{"speed":7.7,"deg":310,"gust":11.3},"clouds":{"all":90},"dt":1572552019,"sys":{"type":1,"id":3551,"country":"US","sunrise":1572525057,"sunset":1572562328},"timezone":-18000,"id":0,"name":"Madison","cod":200}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, mockResp)
	}))
	defer ts.Close()

	fetcher := &mockFetcher{
		fetchFunc: func(ctx context.Context, url string) ([]byte, error) {
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, ts.URL, nil)
			if err != nil {
				return nil, err
			}
			client := &http.Client{}
			res, err := client.Do(req)
			if err != nil {
				return nil, err
			}
			defer res.Body.Close()
			data, err := io.ReadAll(res.Body)
			if err != nil {
				return nil, err
			}
			return data, nil
		},
	}

	ctx := context.Background()
	localWeather, err := GetLocal(ctx, fetcher, 61008, "C")
	if err != nil {
		t.Fatalf("GetLocal failed: %v", err)
	}
	if localWeather.Temp != -0.31 {
		t.Errorf("Temperature is not the correct Celsius temperature: got %f, want -0.31", localWeather.Temp)
	}
}

func TestGetLocalFahrenheit(t *testing.T) {
	mockResp := `{"coord":{"lon":-89.5,"lat":43.03},"weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04d"}],"base":"stations","main":{"temp":272.84,"pressure":1014,"humidity":82,"temp_min":271.48,"temp_max":274.26},"visibility":16093,"wind":{"speed":7.7,"deg":310,"gust":11.3},"clouds":{"all":90},"dt":1572552019,"sys":{"type":1,"id":3551,"country":"US","sunrise":1572525057,"sunset":1572562328},"timezone":-18000,"id":0,"name":"Madison","cod":200}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, mockResp)
	}))
	defer ts.Close()

	fetcher := &mockFetcher{
		fetchFunc: func(ctx context.Context, url string) ([]byte, error) {
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, ts.URL, nil)
			if err != nil {
				return nil, err
			}
			client := &http.Client{}
			res, err := client.Do(req)
			if err != nil {
				return nil, err
			}
			defer res.Body.Close()
			data, err := io.ReadAll(res.Body)
			if err != nil {
				return nil, err
			}
			return data, nil
		},
	}

	ctx := context.Background()
	localWeather, err := GetLocal(ctx, fetcher, 61008, "F")
	if err != nil {
		t.Fatalf("GetLocal failed: %v", err)
	}
	if localWeather.Temp != 31.44 {
		t.Errorf("Temperature is not the correct Fahrenheit temperature: got %f, want 31.44", localWeather.Temp)
	}
}

func TestGetLocalFails(t *testing.T) {
	mockResp := `{"cod":404,"message":"city not found"}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, mockResp)
	}))
	defer ts.Close()

	fetcher := &mockFetcher{
		fetchFunc: func(ctx context.Context, url string) ([]byte, error) {
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, ts.URL, nil)
			if err != nil {
				return nil, err
			}
			client := &http.Client{}
			res, err := client.Do(req)
			if err != nil {
				return nil, err
			}
			defer res.Body.Close()
			data, err := io.ReadAll(res.Body)
			if err != nil {
				return nil, err
			}
			return data, nil
		},
	}

	ctx := context.Background()
	_, err := GetLocal(ctx, fetcher, 6100, "K")
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
