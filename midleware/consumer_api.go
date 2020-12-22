package midleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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
func (*WeatherForeCast) GetByCityAndCountry(city string, country string, apiid string) ForecastResp {
	url := "http://api.openweathermap.org/data/2.5/weather?q=" + city + "," + country + "&units=metric&appid=" + apiid
	resp, er := http.Get(url)
	if er != nil {
		fmt.Println(er.Error())
	}

	forecastData, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		fmt.Println(e.Error())
	}

	var forecastObj Forecast
	json.Unmarshal(forecastData, &forecastObj)

	fmt.Println(forecastObj)

	rsObj := ForecastResp{}
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
