package midleware

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/echenim/globant/models"

	"github.com/echenim/globant/utils/logs"

	"fmt"

	"github.com/echenim/globant/utils"
)

var (
	config utils.IConfigurationManager
)

//NewForeCastAPI psudo contractor
func NewForeCastAPI() IApi {
	config = utils.NewManager()
	return &WeatherForeCast{}
}

//GetByCityAndCountry function defination for fetching forecast
func (*WeatherForeCast) GetByCityAndCountry(city string, country string, apiid string) models.ForecastResp {
	url := "http://api.openweathermap.org/data/2.5/weather?q=" + city + "," + country + "&units=metric&appid=" + apiid
	resp, er := http.Get(url)
	if er != nil {
		l := logs.LogBook{
			Message:  er.Error(),
			Kind:     "FAILED",
			DateTime: time.Now().String()}
		l.Log("API")
	}
	defer resp.Body.Close()
	forecastData, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		l := logs.LogBook{
			Message:  e.Error(),
			Kind:     "FAILED",
			DateTime: time.Now().String()}
		l.Log("API-DATA")
	}

	var forecastObj models.Forecast
	json.Unmarshal(forecastData, &forecastObj)

	rsObj := models.ForecastResp{}
	rsObj.LocationName = forecastObj.Name + " , " + forecastObj.Sys.Country
	rsObj.Temperature = fmt.Sprintf("%.2f", forecastObj.Main.Temp) + " C"
	rsObj.Wind = config.WindScale(forecastObj.Wind.Speed) + ", " + config.Direction(forecastObj.Wind.Deg)
	rsObj.Cloudiness = forecastObj.Weather[0].Description
	rsObj.Pressure = strconv.Itoa(forecastObj.Main.Pressure) + " hpa"
	rsObj.Humidity = strconv.Itoa(forecastObj.Main.Humidity) + "%"
	rsObj.SunRise = config.TimeParse(forecastObj.Sys.SunRise)
	rsObj.SunSet = config.TimeParse(forecastObj.Sys.SunSet)
	rsObj.Coordinates = []float64{forecastObj.Coord.Lon, forecastObj.Coord.Lat}
	rsObj.TimeStamp = config.UnixToHumanReasible(forecastObj.Datetime)

	return rsObj
}

//GetByCityAndCountryAndDay function defination for fetching forecast
func (*WeatherForeCast) GetByCityAndCountryAndDay(city string, country string, apiid string) models.ForecastResp {
	url := "http://api.openweathermap.org/data/2.5/forecast?q=" + city + "," + country + "&" + "&units=metric&appid=" + apiid
	resp, er := http.Get(url)
	if er != nil {
		l := logs.LogBook{
			Message:  er.Error(),
			Kind:     "FAILED",
			DateTime: time.Now().String()}
		l.Log("API")
	}
	defer resp.Body.Close()
	forecastData, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		l := logs.LogBook{
			Message:  e.Error(),
			Kind:     "FAILED",
			DateTime: time.Now().String()}
		l.Log("API-DATA")
	}

	var forecastObj models.Forecast
	json.Unmarshal(forecastData, &forecastObj)

	rsObj := models.ForecastResp{}
	// rsObj.LocationName = forecastObj.Name + " , " + forecastObj.Sys.Country
	// rsObj.Temperature = fmt.Sprintf("%.2f", forecastObj.Main.Temp) + " C"
	// rsObj.Wind = config.WindScale(forecastObj.Wind.Speed) + ", " + config.Direction(forecastObj.Wind.Deg)
	// rsObj.Cloudiness = forecastObj.Weather[0].Description
	// rsObj.Pressure = strconv.Itoa(forecastObj.Main.Pressure) + " hpa"
	// rsObj.Humidity = strconv.Itoa(forecastObj.Main.Humidity) + "%"
	// rsObj.SunRise = config.TimeParse(forecastObj.Sys.SunRise)
	// rsObj.SunSet = config.TimeParse(forecastObj.Sys.SunSet)
	// rsObj.Coordinates = []float64{forecastObj.Coord.Lon, forecastObj.Coord.Lat}
	// rsObj.TimeStamp = config.UnixToHumanReasible(forecastObj.Datetime)

	return rsObj
}
