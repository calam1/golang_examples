package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

// for built in http routing for golang this should work
//func Logger(inner http.Handler, name string) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		start := time.Now()
//
//		inner.ServeHTTP(w, r)
//
//		log.Printf(
//			"%s\t%s\t%s\t%s",
//			r.Method,
//			r.RequestURI,
//			name,
//			time.Since(start),
//		)
//	})
//}

func Logger(fnHandler func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
		start := time.Now()

		fnHandler(w, r, param)

		log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	}
}
