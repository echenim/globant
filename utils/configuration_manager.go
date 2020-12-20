package utils

import (
	"os"
)

//ConfigurationManager function
func ConfigurationManager() *Configurations {
	config := Configurations{}

	configmanager, er := os.Open("appsetting.json")
	if er != nil {

	}

	defer configmanager.Close()

	return &config
}
