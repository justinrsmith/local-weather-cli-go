package main

import (
	"bytes"
	"flag"
	"os"
	"testing"
)

func TestFloatToString(t *testing.T) {
	floatVal := 46.3567
	want := "46.4"
	got := floatToString(floatVal)

	if got != want {
		t.Errorf("Value returned does not match expected")
	}
}

func TestRunNoScale(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"local-weather", "--zipcode", "61008"}
	execute()
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}

func TestRunWithScale(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"local-weather", "--zipcode", "61008", "--scale", "K"}
	execute()
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}

func TestGenerateOutput(t *testing.T) {
	data := []string{
		"Belvidere",
		"Clouds",
		"33.8°F",
		"83%",
		"36.0°F",
		"31.0°F",
	}
	var buf bytes.Buffer
	generateOutput(&buf, data)

	want := `+-----------+-------------------+--------------+----------+-----------+----------+
|   CITY    | CURRENT CONDITION | CURRENT TEMP | HUMIDITY | HIGH TEMP | LOW TEMP |
+-----------+-------------------+--------------+----------+-----------+----------+
| Belvidere | Clouds            | 33.8°F       | 83%      | 36.0°F    | 31.0°F   |
+-----------+-------------------+--------------+----------+-----------+----------+
`
	if buf.String() != want {
		t.Errorf("Generated output does not match expected")
	}
}

func TestGetTempScaleLabelKelvin(t *testing.T) {
	expected := "K"
	got := getTempScaleLabel("K")

	if got != expected {
		t.Errorf("Proper temperature scale label not returned for Kelvin")
	}
}

func TestGetTempScaleLabelFahrenheit(t *testing.T) {
	expected := "\u00b0F"
	got := getTempScaleLabel("F")

	if got != expected {
		t.Errorf("Proper temperature scale label not returned for Fahrenheit")
	}
}

func TestGetTempScaleLabelCelsius(t *testing.T) {
	expected := "\u00b0C"
	got := getTempScaleLabel("C")

	if got != expected {
		t.Errorf("Proper temperature scale label not returned for Celsius")
	}
}

func TestStringInSlice(t *testing.T) {
	expected := true
	got := stringInSlice("K", []string{"C", "K", "F"})

	if got != expected {
		t.Errorf("Expected value not found in slice")
	}
}

func TestStringInSliceNotThere(t *testing.T) {
	expected := false
	got := stringInSlice("Z", []string{"C", "K", "F"})

	if got != expected {
		t.Errorf("Value found in slice when not expected to be there")
	}
}
