package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/echenim/globant/utils"
)

var (
	config *utils.Configurations
)

//NewWeatherController psudo contrustor
func NewWeatherController() IController {

	return &Controller{}
}

//Index function for landing or home page
func (*Controller) Index(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode("hello world")
	return
}

//Get function for fetching records
func (*Controller) Get(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	// _city := chi.URLParam(req, "city")
	// _country := chi.URLParam(req, "country")
	// _day := chi.URLParam(req, "day")

}
