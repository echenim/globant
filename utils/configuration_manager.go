package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

//NewManager psudo constructor
func NewManager() IConfigurationManager {

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

//UnixToHumanReasible function
func (*ConfigManager) UnixToHumanReasible(unixTime int64) string {
	t := time.Unix(unixTime, 0)
	return t.Format(time.UnixDate)
}

//WindScale function
func (*ConfigManager) WindScale(speed float64) string {
	scale := ""
	if speed < 1 {
		scale = "Calm breeze, " + fmt.Sprintf("%.2f", speed) + " m/s"
	}
	if speed >= 1 || speed <= 3 {
		scale = "Light Air, " + fmt.Sprintf("%.2f", speed) + " m/s"
	}
	if speed >= 4 || speed <= 7 {
		scale = "Light Breeze, " + fmt.Sprintf("%.2f", speed) + " m/s"
	}
	if speed >= 8 || speed <= 12 {
		scale = "Gentle Breeze, " + fmt.Sprintf("%.2f", speed) + " m/s"
	}
	if speed >= 13 || speed <= 18 {
		scale = "Moderate, " + fmt.Sprintf("%.2f", speed) + " m/s"
	}
	if speed >= 19 || speed <= 24 {
		scale = "Fresh, " + fmt.Sprintf("%.2f", speed) + " m/s"
	}
	if speed >= 25 || speed <= 31 {
		scale = "Strong, " + fmt.Sprintf("%.2f", speed) + " m/s"
	}

	return scale
}
