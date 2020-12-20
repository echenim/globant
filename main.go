package main

import (
	"fmt"

	"github.com/echenim/globant/handlers"
	"github.com/echenim/globant/utils"
)

var (
	_routes handlers.IRoute
)

func main() {
	config := utils.ConfigurationManager()

	fmt.Println(config)
	// r := chi.NewRouter()

	// _routes = handlers.NewMux(r, config)

	// _routes.SERVE()
}
