package main

import (
	"bytes"
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

func TestRun(t *testing.T) {
	args := os.Args[0:1]
	args = append(args, "--zipcode", "61008") // Append flag
	execute()
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
