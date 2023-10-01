package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/anthonyhungnguyen276/weather-app/types"
	"github.com/anthonyhungnguyen276/weather-app/utils"
)

func GetLatLong(city string) (*types.LatLong, error) {
	endpoint := fmt.Sprintf(utils.GetEnv("LATLONG_API_URL"), url.QueryEscape(city))
	resp, err := http.Get(endpoint)

	if err != nil {
		// %w is used for printing error wrapper message, which provides additional context
		// %v only prints internal error message
		return nil, fmt.Errorf("error making request to GeoAPI: %w", err)
	}
	defer resp.Body.Close()

	// Decode results
	var response types.GeoResponse

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding body: %w", err)
	}

	if len(response.Results) < 1 {
		return nil, fmt.Errorf("no results found")
	}

	return &response.Results[0], nil
}

func GetWeather(latLong *types.LatLong) (string, error) {
	endpoint := fmt.Sprintf(utils.GetEnv("WEATHER_API_URL"), latLong.Latitude, latLong.Longitude)
	resp, err := http.Get(endpoint)

	if err != nil {
		return "", fmt.Errorf("cannot making request to WeatherAPI: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("cannot read weather response: %w", err)
	}

	return string(body), nil

}
