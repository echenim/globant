package controllers

import (
	"net/http"
)

//IController contract defination
type IController interface {
	Get(resp http.ResponseWriter, req *http.Request)
	Find(resp http.ResponseWriter, req *http.Request)
	Index(resp http.ResponseWriter, req *http.Request)
}
