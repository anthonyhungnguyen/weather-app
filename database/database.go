package database

import (
	"github.com/anthonyhungnguyen276/weather-app/types"
	"github.com/jmoiron/sqlx"
)

func InsertCity(db *sqlx.DB, name string, latlong *types.LatLong) error {
	_, err := db.Exec("INSERT INTO cities (name, lat, long) VALUES ($1, $2, $3)", name, latlong.Latitude, latlong.Longitude)
	return err
}
