package v1

import (
	"regexp"
	"math"
	"strconv"
	"strings"
	"fmt"
	dpdts_models "github.com/Logibox/civil-service-interview/v1/bpdts/models"
)

type Location struct {
	Latitude float64
	Longitude float64
}

// Get the distance between two locations in metres using the Haversine formula
// https://en.wikipedia.org/wiki/Haversine_formula
func Distance(loc1, loc2 Location) float64 {
	const r float64 = 6371000.0 // radius of the earth in metres

	hav := func(t float64) float64 {
		return math.Pow(math.Sin(t / 2), 2)
	}

	toRad := func(d float64) float64 {
		return float64(math.Pi * d / 180)
	}

	radlat1 := toRad(loc1.Latitude)
	radlat2 := toRad(loc2.Latitude)
	radlong1 := toRad(loc1.Longitude)
	radlong2 := toRad(loc2.Longitude)

	h := hav(radlat2 - radlat1) + math.Cos(radlat1) * math.Cos(radlat2) * hav(radlong2 - radlong1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

func MetresToMiles(metres float64) float64 {
	return metres * 0.00062137
}

func UserLocation(user *dpdts_models.User) (*Location, error) {
	lat, err := user.Latitude.Float64()
	if err != nil {
		return nil, err
	}
	long, err := user.Longitude.Float64()
	if err != nil {
		return nil, err
	}
	return &Location{
		Latitude: lat,
		Longitude: long,
	}, nil
}

// Parse a string and return the distance in metres
// Examples:
// ParseDistanceString("50m") -> 50.0
// ParseDistanceString("50km") -> 50000.0
// ParseDistanceString("50 miles") -> 80467.2
func ParseDistanceString(distanceStr string) (float64, error) {
	re := regexp.MustCompile(`(\d+)\s*(\w+)`)
	matches := re.FindAllStringSubmatch(distanceStr, -1)
	if len(matches) != 1 || len(matches[0]) != 3 {
		return 0.0, fmt.Errorf("invalid distance string format: %s", distanceStr)
	}

	d, err := strconv.ParseFloat(matches[0][1], 64)
	if err != nil {
		return 0.0, err
	}

	switch strings.ToLower(matches[0][2]) {
	case "metres":
		fallthrough
	case "meters":
		fallthrough
	case "m":
		return d, nil
	case "miles":
		return d * 1609.34, nil
	case "kilometres":
		fallthrough
	case "kilometers":
		fallthrough
	case "km":
		return d * 1000.0, nil
	default:
		return 0.0, fmt.Errorf("invalid distance unit: %s", matches[0][2])
	}
}
