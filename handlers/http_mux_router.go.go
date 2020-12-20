package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

var (
	dispatcher *chi.Mux
)

//NewMux is a psudo constractor
func NewMux(_dispatcher *chi.Mux) IRoute {
	dispatcher = _dispatcher
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	dispatcher.Get(uri, f)
}
func (*muxRouter) SERVE(port string) {
	http.ListenAndServe(":"+port, dispatcher)
}
