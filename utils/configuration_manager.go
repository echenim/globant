package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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
	year, month, day := t.Date()
	return month.String() + " " + strconv.Itoa(day) + ", " + strconv.Itoa(year) + " " + strconv.Itoa(t.Hour()) + ":" + strconv.Itoa(t.Minute()) + ":" + strconv.Itoa(t.Second())
}

//TimeParse function
func (*ConfigManager) TimeParse(unixTime int64) string {
	t := time.Unix(unixTime, 0)
	return ampmformatta(t.Hour(), t.Minute())
}

func ampmformatta(h int, m int) string {
	if h < 12 {
		return strconv.Itoa(h) + ":" + strconv.Itoa(m) + "AM"
	}
	hh := h - 12
	return strconv.Itoa(hh) + ":" + strconv.Itoa(m) + "PM"
}

//WindScale function
func (*ConfigManager) WindScale(speed float64) string {
	scale := ""
	if speed < 1.00 {
		scale = "Calm breeze, " + fmt.Sprintf("%.2f", speed) + " m/s"
	}
	if speed >= 1.00 && speed <= 3.00 {
		scale = "Light Air, " + fmt.Sprintf("%.2f", speed) + " m/s"
	}
	if speed > 3.00 && speed <= 7.00 {
		scale = "Light Breeze, " + fmt.Sprintf("%.2f", speed) + " m/s"
	}
	if speed > 7.00 && speed <= 12.00 {
		scale = "Gentle Breeze, " + fmt.Sprintf("%.2f", speed) + " m/s"
	}
	if speed > 12.00 && speed <= 18.00 {
		scale = "Moderate, " + fmt.Sprintf("%.2f", speed) + " m/s"
	}
	if speed > 18.00 && speed <= 24.00 {
		scale = "Fresh, " + fmt.Sprintf("%.2f", speed) + " m/s"
	}
	if speed > 24.00 && speed <= 31.00 {
		scale = "Strong, " + fmt.Sprintf("%.2f", speed) + " m/s"
	}

	if speed > 31.00 && speed <= 38.00 {
		scale = "Near Gale, " + fmt.Sprintf("%.2f", speed) + " m/s"
	}

	if speed > 38.00 && speed <= 46.00 {
		scale = "Gale, " + fmt.Sprintf("%.2f", speed) + " m/s"
	}

	if speed > 46.00 && speed <= 54.00 {
		scale = "Strong Gale, " + fmt.Sprintf("%.2f", speed) + " m/s"
	}

	if speed > 54.00 && speed <= 63.00 {
		scale = "Storm, " + fmt.Sprintf("%.2f", speed) + " m/s"
	}
	if speed > 63.00 && speed <= 73.00 {
		scale = "Violent Storm, " + fmt.Sprintf("%.2f", speed) + " m/s"
	}
	if speed > 73.00 && speed <= 95.00 {
		scale = "Hurricane, " + fmt.Sprintf("%.2f", speed) + " m/s"
	}
	//fmt.Println("wind speed ", speed, "   ", scale)
	return scale
}

//Direction function converter
func (*ConfigManager) Direction(deg float64) string {
	scale := ""

	if deg == 0.0 {
		scale = " north "
	}
	if deg > 0.0 && deg < 90.00 {
		scale = "north-east "
	}
	if deg == 90.00 {
		scale = " east "
	}
	if deg > 90.00 && deg < 180.00 {
		scale = "south-east "
	}
	if deg == 180.00 {
		scale = " south "
	}
	if deg > 180.00 && deg < 270.00 {
		scale = "south-west "
	}
	if deg == 270.00 {
		scale = " west "
	}
	if deg > 270.00 && deg <= 360.00 {
		scale = "north-west "
	}

	return scale
}

//Cloudiness function
func (*ConfigManager) Cloudiness(cloudiness float64) string {
	scale := ""
	if cloudiness < 6 {
		scale = " Clear"
	}
	if cloudiness >= 6 && cloudiness < 26 {
		scale = "  Mostly Clear"
	}
	if cloudiness >= 26 && cloudiness < 51 {
		scale = " Partly Clear"
	}
	if cloudiness >= 51 && cloudiness < 70 {
		scale = " Mostly Clear"
	}
	if cloudiness >= 88 && cloudiness <= 100 {
		scale = " Cloudy"
	}

	return scale
}
