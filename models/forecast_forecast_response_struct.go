package models

//ForecastResp response defination
type ForecastResp struct {
	LocationName string    `json:"location_name"`
	Temperature  string    `json:"temperature"`
	Wind         string    `json:"wind"`
	Cloudiness   string    `json:"cloudiness"`
	Pressure     string    `json:"pressure"`
	Humidity     string    `json:"humidity"`
	SunRise      string    `json:"sunrise"`
	SunSet       string    `json:"sunset"`
	Coordinates  []float64 `json:"geo_coordinates"`
	TimeStamp    string    `json:"requested_time"`
}

