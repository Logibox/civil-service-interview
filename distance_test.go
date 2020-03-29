package v1_test

import (
	"testing"
	csi "github.com/Logibox/civil-service-interview/v1"
)

func TestDistance(t *testing.T) {
	manchesterLoc := csi.Location{53.4808, -2.2426}
	londonLoc := csi.Location{51.5074, -0.1278}

	distance := csi.Distance(manchesterLoc, londonLoc)
	if int(distance) != 261982 {
		t.Errorf("Distance from Manchester to London is not %v metres", distance) 
	}
}

func TestParseDistanceString(t *testing.T) {
	data := []struct {
		input string
		output int64
	}{
		{"50 meters", 50},
		{"20 miles", 32186},
		{"5m", 5},
		{"100    km", 100000},
		{"15 kilometres", 15000},
	}

	for _, d := range data {
		o, err := csi.ParseDistanceString(d.input)
		if err != nil {
			t.Errorf("error parsing %s: %v", d.input, err)
		}
		if int64(o) != d.output {
			t.Errorf("error parsing %s: got %d expected %d", d.input, int64(o), d.output)
		}
	}

	badData := []string {
		"not a distance",
		"50",
		"13 apples",
	}
	for _, d := range badData {
		_, err := csi.ParseDistanceString(d)
		if err == nil {
			t.Errorf("error string %s should report an error", d)
		}
	}
}
