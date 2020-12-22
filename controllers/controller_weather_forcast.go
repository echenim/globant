package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/echenim/globant/midleware"
	"github.com/echenim/globant/models"
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
	_country := strings.ToLower(chi.URLParam(req, "country"))

	resultset := models.ForecastResp{}

	if _city != "" && _country != "" {
		resultset = _warz.GetByCityAndCountry(_city, _country, apikey)
		resp.WriteHeader(http.StatusOK)
		json.NewEncoder(resp).Encode(resultset)

	} else {
		resp.WriteHeader(http.StatusBadRequest)
	}

	return
}

//Find function for fetching records
func (*Controller) Find(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	_city := chi.URLParam(req, "city")
	_country := strings.ToLower(chi.URLParam(req, "country"))
	//_day := chi.URLParam(req, "day")

	resultset := models.SixDaysForecastSetResp{}

	if _city != "" && _country != "" {
		resultset = _warz.GetByCityAndCountryAndDay(_city, _country, apikey)
		resp.WriteHeader(http.StatusOK)
		json.NewEncoder(resp).Encode(resultset)

	} else {
		resp.WriteHeader(http.StatusBadRequest)
	}
	return
}

