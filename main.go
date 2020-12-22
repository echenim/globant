package main

import (
	"fmt"
	"time"

	"github.com/echenim/globant/controllers"
	"github.com/echenim/globant/midleware"

	"github.com/echenim/globant/handlers"
	"github.com/echenim/globant/utils"
)

var (
	_routes     handlers.IRoute
	_config     utils.IConfigurationManager
	_controller controllers.IController
	_warz       midleware.IApi
)

func main() {
	_config = utils.NewManager()
	_warz = midleware.NewForeCastAPI()
	// _controller = controllers.NewWeatherController()
	// r := chi.NewRouter()

	// _routes = handlers.NewMux(r)
	// _routes.GET("/", _controller.Index)
	// _routes.GET("/weatherforecast/{city},{country}", _controller.Get)
	// _routes.GET("/weatherforecast/{city},{country}/{day}", _controller.Get)

	// _routes.SERVE(_config.ConfigurationManager().ConfigurationManagerList[0].Port)

	// k := _warz.GetByCityAndCountry("Shuzenji", "jp", _config.ConfigurationManager().ConfigurationManagerList[0].APIKey)
	// fmt.Println(k)

	t := time.Unix(1608587328, 0)
	fmt.Println("hello ", t.Hour())
}
