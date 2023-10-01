package types

// The json tags tell the JSON decoder how to map the JSON fields to the struct fields. Extra fields in the JSON response are ignored by default.

type LatLong struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GeoResponse struct {
	Results []LatLong `json:"results"`
}
