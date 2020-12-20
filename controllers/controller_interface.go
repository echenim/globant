package controllers

import (
	"net/http"
)

type IController interface {
	Get(resp http.ResponseWriter, req *http.Request)
	Index(resp http.ResponseWriter, req *http.Request)
}
