package midleware

//WeatherForeCast defination from API
type WeatherForeCast struct {
}

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

//Forecast defination
type Forecast struct {
	Name     string      `json:"name"`
	Coord    Coordinates `json:"coord"`
	Weather  []Weathers  `json:"weather"`
	Base     string      `json:"base"`
	Main     Mains       `json:"main"`
	Wind     Winds       `json:"wind"`
	Cloud    Clouds      `json:"clouds"`
	Sys      System      `json:"sys"`
	Datetime int64      `json:"dt"`
}

//Coordinates defination
type Coordinates struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

//Weathers defination
type Weathers struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

//Mains defination
type Mains struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	Tempmax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

//Winds defination
type Winds struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}

//Clouds defination
type Clouds struct {
	All int `json:"all"`
}

//System defination
type System struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	SunRise int64  `json:"sunrise"`
	SunSet  int64  `json:"sunset"`
}
