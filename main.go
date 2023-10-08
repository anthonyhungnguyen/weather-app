package main

import (
	"log"
	"net/http"

	"github.com/anthonyhungnguyen276/weather-app/api"
	"github.com/anthonyhungnguyen276/weather-app/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func QueryWeatherByCity(db *sqlx.DB, city string, c *gin.Context) {
	latlong, err := api.GetLatLong(db, city)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	weather, err := api.GetWeather(latlong)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	weatherDisplay, err := utils.ExtractWeatherData(city, weather)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"weather": weatherDisplay})
}

func GetLastSearchedCities(db *sqlx.DB) ([]string, error) {
	var cities []string
	err := db.Select(&cities, "SELECT name FROM cities ORDER BY id DESC LIMIT 10")
	if err != nil {
		return nil, err
	}
	return cities, nil
}

func main() {

	r := gin.Default()
	db, err := sqlx.Connect("postgres", utils.GetEnv("DATABASE_URL"))

	if err != nil {
		log.Fatalf("error connecting to database: %w", err)
	}

	r.GET("/weather", func(c *gin.Context) {
		city := c.Query("city")
		QueryWeatherByCity(db, city, c)
	})

	r.GET("/weather/:city", func(c *gin.Context) {
		city := c.Param("city")
		QueryWeatherByCity(db, city, c)
	})

	r.GET("/stats", gin.BasicAuth(gin.Accounts{
		"forecast": "forecast",
	}), func(c *gin.Context) {
		cities, err := GetLastSearchedCities(db)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"stats": cities})
	})

	r.Run()
}
