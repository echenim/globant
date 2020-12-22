package models

//SixDaysForecastList defination struct
type SixDaysForecastList struct {
	ForecastDate string `json:"forecast_date"`
	Temperature  string `json:"temperature"`
	Wind         string `json:"wind"`
	Cloudiness   string `json:"cloudiness"`
	Pressure     string `json:"pressure"`
	Humidity     string `json:"humidity"`
	TimeStamp    string `json:"requested_time"`
}