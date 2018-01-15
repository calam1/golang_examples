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
		"/files",
		FilesInfo,
	},
	Route{
		"GET",
		"/files/file/:name",
		FileContent,
	},
}
