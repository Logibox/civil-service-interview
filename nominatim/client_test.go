package nominatim_test

import (
	"fmt"
	"testing"
	"github.com/Logibox/civil-service-interview/v1/nominatim"
)

func TestGetCityLocation(t *testing.T) {
	ld, err := nominatim.GetCityLocation("Manchester", "UK")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	fmt.Println(ld)
}
