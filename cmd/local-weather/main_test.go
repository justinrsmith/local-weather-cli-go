package main

import (
	"bytes"
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

func TestGenerateOutput(t *testing.T) {
	data := []string{
		"Belvidere",
		"Clear",
		"68.9°F",
		"82%",
		"3.2mph West-Northwest",
		"69.3°F",
	}
	var buf bytes.Buffer
	generateOutput(&buf, data)

	want := `+-----------+-------------------+--------------+----------+-----------------------+------------+
|   CITY    | CURRENT CONDITION | CURRENT TEMP | HUMIDITY |         WIND          | FEELS LIKE |
+-----------+-------------------+--------------+----------+-----------------------+------------+
| Belvidere | Clear             | 68.9°F       | 82%      | 3.2mph West-Northwest | 69.3°F     |
+-----------+-------------------+--------------+----------+-----------------------+------------+
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
