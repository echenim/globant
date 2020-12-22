package main

import (
	"github.com/echenim/globant/controllers"
	"github.com/go-chi/chi"

	"github.com/echenim/globant/handlers"
	"github.com/echenim/globant/utils"
)

var (
	_routes     handlers.IRoute
	_config     utils.IConfigurationManager
	_controller controllers.IController
)

func main() {
	_config = utils.NewManager()
	apiID := _config.ConfigurationManager().ConfigurationManagerList[0].APIKey
	_controller = controllers.NewWeatherController(apiID)
	r := chi.NewRouter()

	_routes = handlers.NewMux(r)
	_routes.GET("/", _controller.Index)
	_routes.GET("/weatherforecast/{city},{country}", _controller.Get)
	_routes.GET("/weatherforecast/{city},{country}/{day}", _controller.Find)
	_routes.SERVE(_config.ConfigurationManager().ConfigurationManagerList[0].Port)

}
