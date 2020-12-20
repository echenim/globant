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
func NewWeatherController(_config *utils.Configurations) IController {
	config = _config

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

}
