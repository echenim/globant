package models

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
	Datetime int64       `json:"dt"`
}
