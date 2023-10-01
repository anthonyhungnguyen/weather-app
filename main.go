package main

import (
	"fmt"
	"log"

	"github.com/anthonyhungnguyen276/weather-app/api"
)

func main() {
	city := "Windsor"
	latlong, err := api.GetLatLong(city)

	if err != nil {
		log.Fatalf("failed to get latlong: %s", err)
	}

	fmt.Printf("Latitude %f, Longitude: %f\n", latlong.Latitude, latlong.Longitude)

	weather, err := api.GetWeather(latlong)

	if err != nil {
		log.Fatalf("failed to get weather: %s\n", err)
	}

	fmt.Printf("Weather: %s", weather)
}
