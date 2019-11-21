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
