package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os/exec"
)

func RunJob(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	directory := request.FormValue("directory")
	jobName := params.ByName("name")

	fullPath := directory + "/" + jobName
	cmdOut := exec.Command("sh", fullPath)

	cmdOutput, err := cmdOut.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(cmdOutput))

}
