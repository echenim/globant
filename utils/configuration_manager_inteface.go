package utils

//IConfigurationManager contract defination
type IConfigurationManager interface {
	ConfigurationManager() *ConfigurationList
	UnixToHumanReasible(unixTime int64) string
	WindScale(speed float64) string
}
