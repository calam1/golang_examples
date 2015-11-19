package main

import (
	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	for _, route := range routes {
		var handler httprouter.Handle
		handler = route.Handle
		handler = Logger(handler)

		router.Handle(route.Method, route.Path, handler)
	}
	return router
}
