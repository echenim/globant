package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/echenim/globant/midleware"
	"github.com/go-chi/chi"
)

var (
	_warz  midleware.IApi
	apikey string
)

//NewWeatherController psudo contrustor
func NewWeatherController(key string) IController {
	_warz = midleware.NewForeCastAPI()
	apikey = key
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
	_city := chi.URLParam(req, "city")
	_country := chi.URLParam(req, "country")
	_day := chi.URLParam(req, "day")

	resultset := midleware.ForecastResp{}

	fmt.Println(_country)
	if _city != "" && _country != "" && _day != "" {
		//resultset = _warz.GetByCityAndCountry(_city, _country, _config.ConfigurationManager().ConfigurationManagerList[0].APIKey)
	}

	if _city != "" && _country != "" {
		resultset = _warz.GetByCityAndCountry(_city, _country, apikey)
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(resultset)

	return
}
