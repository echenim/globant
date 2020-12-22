package models

//SixDaysForecastSetRespList defination struct
type SixDaysForecastSetRespList struct {
	LocationName string                `json:"location_name"`
	SunRise      string                `json:"sunrise"`
	SunSet       string                `json:"sunset"`
	Coordinates  []float64             `json:"geo_coordinates"`
	Forecast     []SixDaysForecastList `json:"forecast_list"`
}


