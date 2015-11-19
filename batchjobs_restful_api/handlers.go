package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func JobsInfo(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	//we can take a commmand line argument if one is passed
	flag.Parse()
	dirRoot := flag.Arg(0)

	//if url is passed this will override the passed argument at start up
	directory := request.FormValue("directory")
	prettyPrint := request.FormValue("prettyprint")

	if directory != "" {
		dirRoot = directory
	}

	jobNames := getJobNames(dirRoot)
	jobs := Jobs{jobNames}

	var jobsJson []byte
	var err error
	if prettyPrint == "true" {
		jobsJson, err = json.MarshalIndent(jobs, "", "    ")
	} else {
		jobsJson, err = json.Marshal(jobs)
	}

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jobsJson)
}

//use of closure and anonymous function
func getJobNames(directory string) []JobName {
	jobNames := make([]JobName, 0)
	err := filepath.Walk(directory, func(dir string, f os.FileInfo, err error) error {
		matched, err := filepath.Match("*.sh", f.Name())

		if err != nil {
			fmt.Println(err)
		}

		if matched {
			dirName, fileName := filepath.Split(dir)
			contentUrl := "/jobs/job/" + fileName + "?directory=" + dirName
			jobName := JobName{fileName, dirName, contentUrl, "/fakeurl/execute"}
			jobNames = append(jobNames, jobName)
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return jobNames
}

func JobContent(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	directory := request.FormValue("directory")
	fileName := params.ByName("name")
	fullPath := directory + "/" + fileName

	data, err := ioutil.ReadFile(fullPath)
	if err != nil {
		panic(err)
	}

	writer.Header().Set("Content-Type", "text/plain")
	writer.WriteHeader(http.StatusOK)
	writer.Write(data)
}
