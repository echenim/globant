package midleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Fetch function
func Fetch() {
	resp, er := http.Get("http://api.openweathermap.org/data/2.5/weather?q=Bogota,co&appid=1508a9a4840a5574c822d70ca2132032")
	if er != nil {
		fmt.Println(er.Error())
	}

	respData, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		fmt.Println(e.Error())
	}

	var rsObj Forecast
	json.Unmarshal(respData, &rsObj)
	fmt.Println(rsObj)
}

//NewForeCastAPI psudo contractor
func NewForeCastAPI() IApi {
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
	rsObj.Wind = 

	return rsObj
}
