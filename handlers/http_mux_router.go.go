package handlers

import (
	"net/http"

	"github.com/echenim/globant/utils"

	"github.com/go-chi/chi"
)

var (
	dispatcher *chi.Mux
	port       string
)

//NewMux is a psudo constractor
func NewMux(_dispatcher *chi.Mux, _config *utils.Configurations) IRoute {
	dispatcher = _dispatcher
	port = _config.Port
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	dispatcher.Get(uri, f)
}
func (*muxRouter) SERVE() {
	http.ListenAndServe(":"+port, dispatcher)
}
