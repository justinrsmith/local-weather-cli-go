package main

import (
	"bytes"
	"fmt"
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

func TestKelvinToFarhenheit(t *testing.T) {
	got := kelvinToFarhenheit(280.040000)
	fmt.Println(got)
	want := 44.40199999999999
	if got != want {
		t.Errorf("No match")
	}
}

func TestRun(t *testing.T) {
	args := os.Args[0:1]
	args = append(args, "--zipcode", "61008") // Append flag
	run(args)
}

func TestGenerateOutput(t *testing.T) {
	data := []string{
		"Belvidere",
		"Sunny",
		"52.5",
		"85%",
		"60.2",
		"40.5",
	}
	var buf bytes.Buffer
	generateOutput(&buf, data)

	want := `+-----------+-------------------+--------------+----------+-----------+----------+
|   CITY    | CURRENT CONDITION | CURRENT TEMP | HUMIDITY | HIGH TEMP | LOW TEMP |
+-----------+-------------------+--------------+----------+-----------+----------+
| Belvidere | Sunny             |         52.5 | 85%      |      60.2 |     40.5 |
+-----------+-------------------+--------------+----------+-----------+----------+
`
	if buf.String() != want {
		t.Errorf("Generated output does not match expected")
	}
}
