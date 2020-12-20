package main

import (
	"github.com/echenim/globant/handlers"
	"github.com/echenim/globant/utils"
	"github.com/go-chi/chi"
)

var (
	_routes handlers.IRoute
)

func main() {
	config := utils.ReadConfig()
	r := chi.NewRouter()

	_routes = handlers.NewMux(r, config)

	_routes.SERVE()
}
