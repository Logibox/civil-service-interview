package nominatim

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// [{"place_id":100032,"licence":"Data © OpenStreetMap contributors, ODbL 1.0. https://osm.org/copyright","osm_type":"node","osm_id":107775,"boundingbox":["51.3473219","51.6673219","-0.2876474","0.0323526"],"lat":"51.5073219","lon":"-0.1276474","display_name":"London, Greater London, England, SW1A 2DX, United Kingdom","class":"place","type":"city","importance":0.9407827616237295,"icon":"https://nominatim.openstreetmap.org/images/mapicons/poi_place_city.p.20.png"}]

type LocationData struct {
	Latitude json.Number `json:"lat"`
	Longitude json.Number `json:"lon"`
	DisplayName string `json:"display_name"`
}

const nominatimURL = "https://nominatim.openstreetmap.org/search/"

// const https://nominatim.openstreetmap.org/search/?format=json&city=London&country=UK

var (
	cache map[string]*LocationData
	rateLimiter *time.Ticker
)

func init() {
	cache = make(map[string]*LocationData)
	// The terms of uses dictates 1 request per second, but lets play it safe
	rateLimiter = time.NewTicker(2 * time.Second)
}

func GetCityLocation(city string, country string) (*LocationData, error) {
	query := fmt.Sprintf("?format=json&city=%s&country=%s", city, country)
	if ld, ok := cache[query]; ok {
		// Return cached result
		return ld, nil
	}

	<-rateLimiter.C
	resp, err := http.Get(nominatimURL + query)
	if err != nil {
		return nil, err
	}

	arr := make([]LocationData, 1)
	ld := &arr[0]
	if err = json.NewDecoder(resp.Body).Decode(&arr); err != nil {
		return nil, err
	}
	cache[query] = ld
	return ld, nil
}
