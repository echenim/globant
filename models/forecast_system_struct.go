package models

//System defination
type System struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	SunRise int64  `json:"sunrise"`
	SunSet  int64  `json:"sunset"`
}