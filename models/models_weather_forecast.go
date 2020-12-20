package models

import (
	"time"
)

//Forecast models struct
type Forecast struct {
	LocationName string
	Temperature  string
	WindStatus string
	Cloudiness string
	Pressure string
	Sunrise string
	Coordinates string
	RequestTime time.Time
	
}
