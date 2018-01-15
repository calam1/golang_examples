package main

import (
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func FileContent(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	directory := request.FormValue("directory")
	fileName := params.ByName("name")
	fullPath := directory + "/" + fileName

	//for now we just read the whole file into memory
	data, err := ioutil.ReadFile(fullPath)
	if err != nil {
		panic(err)
	}

	writer.Header().Set("Content-Type", "text/plain")
	writer.WriteHeader(http.StatusOK)
	writer.Write(data)
}
