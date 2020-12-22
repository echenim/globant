package utils

//IConfigurationManager contract defination
type IConfigurationManager interface {
	ConfigurationManager() *ConfigurationList
	UnixToHumanReasible(unixTime int64) string
	WindScale(speed float64) string
	Direction(deg float64) string
	TimeParse(unixTime int64) string
}
