package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/anthonyhungnguyen276/weather-app/types"
	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("cannot read dotenv file: %w", err)
	}

	return os.Getenv(key)
}

func ExtractWeatherData(city string, rawWeather string) (types.WeatherDisplay, error) {
	var weatherResponse types.WeatherResponse
	if err := json.Unmarshal([]byte(rawWeather), &weatherResponse); err != nil {
		return types.WeatherDisplay{}, fmt.Errorf("error decoding weather response %w", err)
	}

	var forecasts []types.Forecast
	for i, t := range weatherResponse.Hourly.Time {
		date, err := time.Parse("2006-01-02T15:04", t)
		if err != nil {
			return types.WeatherDisplay{}, nil
		}
		forecast := types.Forecast{
			Date:        date.Format("Mon 15:04"),
			Temperature: fmt.Sprintf("%1.f°C", weatherResponse.Hourly.Temperature2M[i]),
		}

		forecasts = append(forecasts, forecast)
	}
	return types.WeatherDisplay{
		City:      city,
		Forecasts: forecasts,
	}, nil
}
