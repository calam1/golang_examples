package main

import (
	"net/http"
)

//usage i.e. curl localhost:8990?directory=/home/test/scripts/ if you want to pass in the path to search
//otherwise if you want to use the directory where the golang process is started(if process is started with the . arg)
//and don't add the directory param
//if you want pretty print of the json, add the parameter prettyprint=true to the url
func main() {
	router := NewRouter()
	http.ListenAndServe(":8990", router)
}
