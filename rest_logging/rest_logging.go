package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	//"io/ioutil"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	Trace *log.Logger
	Info  *log.Logger
)

func init() {
	file, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", ":", err)
	}

	multi := io.MultiWriter(file, os.Stdout)

	Trace = log.New(multi,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(multi,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Trace.Println("Start of the spawn_job_process process.")

	Trace.Println("Creating new httprouter to route RESTful paths.")
	router := httprouter.New()

	// Add a handler
	Info.Println("Setting up the handler for GET /test request.")
	router.GET("/test", func(writer http.ResponseWriter, router *http.Request, _ httprouter.Params) {
		Trace.Println("Handling the GET /test RESTful path.")
		fmt.Fprint(writer, "Welcome!\n")
	})

	// Fire up the server
	Info.Println("Starting the server on port 8990.")
	//http.ListenAndServe("localhost:3000", router)
	http.ListenAndServe(":8990", router)
}
