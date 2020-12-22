package models

//SixDaysForecast defination
type SixDaysForecast struct {
	ForecastSet     []SixDaysForecastSet   `json:"list"`
	CityInformation SixDaysCityInformation `json:"city"`
}
