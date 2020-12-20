package handlers

import (
	"net/http"
)

//IRoute interface defination
type IRoute interface {
	GET(uri string, f func(resp http.ResponseWriter, req *http.Request))
	SERVE(port string)
}
