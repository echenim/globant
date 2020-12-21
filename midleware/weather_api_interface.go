package midleware

//IApi contract defination
type IApi interface {
	GetByCityAndCountry(city string, country string, apiid string) ForecastResp
}
