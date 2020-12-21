package main

import (
	"github.com/echenim/globant/controllers"

	"github.com/echenim/globant/handlers"
	"github.com/echenim/globant/utils"
	"github.com/go-chi/chi"
)

var (
	_routes     handlers.IRoute
	_config     utils.IConfigurationManager
	_controller controllers.IController
)

func main() {
	_config = utils.NewConfigurationManager()
	_controller = controllers.NewWeatherController()
	r := chi.NewRouter()

	_routes = handlers.NewMux(r)
	_routes.GET("/", _controller.Index)
	_routes.GET("/weatherforecast/{city},{country}", _controller.Get)
	_routes.GET("/weatherforecast/{city},{country}/{day}", _controller.Get)

	_routes.SERVE(_config.ConfigurationManager().ConfigurationManagerList[0].Port)
}
