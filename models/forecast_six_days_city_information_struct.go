package models

//SixDaysCityInformation defination struct
type SixDaysCityInformation struct {
	Name       string      `json:"name"`
	Coord      Coordinates `json:"coord"`
	Country    string      `json:"country"`
	SunRise    int64       `json:"sunrise"`
	SunSet     int64       `json:"sunset"`
	Population int64       `json:"population"`
}
