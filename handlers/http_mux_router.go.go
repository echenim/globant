package handlers

import "github.com/go-chi/chi"

var (
	dispatcher *chi.Mux
)

//NewMux is a psudo constractor
func NewMux(_dispatcher *chi.Mux) IRoute {
	dispatcher = _dispatcher
	return &muxRouter{}
}
