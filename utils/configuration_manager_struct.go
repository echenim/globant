package utils

//Configurations strunct
type Configurations struct {
	Port          string `json:"port"`
	CacheDuration int    `json:"cacheduration"`
}

//ConfigurationList collection
type ConfigurationList struct {
	ConfigurationManagerList []Configurations `json:"configurationmanagerlist"`
}

//ConfigManager empty struct defination
type ConfigManager struct{}
