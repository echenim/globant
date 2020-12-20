package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//NewConfigurationManager psudo constructor
func NewConfigurationManager() IConfigurationManager {

	return &ConfigManager{}
}

//ConfigurationManager function
func (*ConfigManager) ConfigurationManager() *ConfigurationList {

	configmanager, er := os.Open("appsetting.json")
	if er != nil {
		fmt.Println(er.Error())
	}
	defer configmanager.Close()
	configData, e := ioutil.ReadAll(configmanager)
	if er != nil {
		fmt.Println(e.Error())
	}

	var readConfig *ConfigurationList

	json.Unmarshal(configData, &readConfig)

	return readConfig
}
