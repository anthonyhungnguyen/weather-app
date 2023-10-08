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

	c.HTML(http.StatusOK, "weather.html", weatherDisplay)

	// c.JSON(http.StatusOK, gin.H{"weather": weather})
}

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("views/*")
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

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run()
}
