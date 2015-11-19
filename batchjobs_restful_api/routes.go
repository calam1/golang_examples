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

var routes = Routes{
	Route{
		"GET",
		"/jobs",
		JobsInfo,
	},
	Route{
		"GET",
		"/jobs/job/:name",
		JobContent,
	},
}
