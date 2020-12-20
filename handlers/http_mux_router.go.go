package handlers

import (
	"fmt"
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
	fmt.Println("Server is starting .............")
	fmt.Println("Server is listing on Port :" + port)
	http.ListenAndServe(":"+port, dispatcher)

}
