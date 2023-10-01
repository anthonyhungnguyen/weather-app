package main

import (
	"net/http"

	"github.com/anthonyhungnguyen276/weather-app/api"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/weather", func(c *gin.Context) {
		city := c.Query("city")
		latlong, err := api.GetLatLong(city)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		weather, err := api.GetWeather(latlong)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"weather": weather})
	})

	r.Run()
}
