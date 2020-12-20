package main

import (
	"fmt"

	"github.com/echenim/globant/handlers"
	"github.com/echenim/globant/utils"
	"github.com/go-chi/chi"
)

var (
	_routes handlers.IRoute
	_config utils.IConfigurationManager
)

func main() {
	_config = utils.NewConfigurationManager()

	fmt.Println("hello  ", _config.ConfigurationManager().ConfigurationManagerList[0].Port)
	r := chi.NewRouter()

	_routes = handlers.NewMux(r)

	_routes.SERVE(_config.ConfigurationManager().ConfigurationManagerList[0].Port)
}
