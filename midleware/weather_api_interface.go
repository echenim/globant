package midleware

import "github.com/echenim/globant/models"

//IApi contract defination
type IApi interface {
	GetByCityAndCountry(city string, country string, apiid string) models.ForecastResp
	GetByCityAndCountryAndDay(city string, country string, apiid string) models.SixDaysForecastSetResp
}
