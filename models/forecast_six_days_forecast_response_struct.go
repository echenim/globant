package models

//SixDaysForecastSetResp defination struct
type SixDaysForecastSetResp struct {
	LocationName string    `json:"location_name"`
	SunRise      string    `json:"sunrise"`
	SunSet       string    `json:"sunset"`
	Coordinates  []float64 `json:"geo_coordinates"`
	ForecastDate string    `json:"forecast_date"`
	Temperature  string    `json:"temperature"`
	Wind         string    `json:"wind"`
	Cloudiness   string    `json:"cloudiness"`
	Pressure     string    `json:"pressure"`
	Humidity     string    `json:"humidity"`
	TimeStamp    string    `json:"requested_time"`
}
