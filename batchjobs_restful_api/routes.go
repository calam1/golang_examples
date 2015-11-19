package main

import (
	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Method string
	Path   string
	Handle httprouter.Handle
}

type Routes []Route

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

var routes = Routes{
	Route{
		"GET",
		"/jobnames",
		JobNames,
	},
}
