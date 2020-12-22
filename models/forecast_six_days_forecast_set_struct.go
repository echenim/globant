package models

//SixDaysForecastSet  deination struct
type SixDaysForecastSet struct {
	Datetime        int64                  `json:"dt"`
	Main            Mains                  `json:"main"`
	Weather         []Weathers             `json:"weather"`
	Cloud           Clouds                 `json:"clouds"`
	Wind            Winds                  `json:"wind"`
	Date            string                 `json:"dt_txt"`
	CityInformation SixDaysCityInformation `json:"city"`
}
