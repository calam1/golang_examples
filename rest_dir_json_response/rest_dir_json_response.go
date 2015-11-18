package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"path/filepath"
)

type JobName struct {
	Name string
	Path string
}

func init() {
	fmt.Println("initialize block")
}

//use of closure and anonymous function
func getJobNames(directory string) []JobName {
	list := make([]JobName, 0)
	err := filepath.Walk(directory, func(dir string, f os.FileInfo, err error) error {
		matched, err := filepath.Match("*.sh", f.Name())

		if err != nil {
			fmt.Println(err)
		}

		if matched {
			dirName, fileName := filepath.Split(dir)
			jobName := JobName{fileName, dirName}
			list = append(list, jobName)
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return list
}

//usage i.e. curl localhost:8990?directory=/home/test/scripts/ if you want to pass in the path to search
//otherwise if you want to use the directory where the golang process is started  don't add the directory param
func main() {
	router := httprouter.New()
	router.GET("/jobnames", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		//we can take a commmand line argument if one is passed
		flag.Parse()
		dirRoot := flag.Arg(0)

		//if url is passed this will override the passed argument at start up
		directory := request.FormValue("directory")
		if directory != "" {
			dirRoot = directory
		}

		list := getJobNames(dirRoot)
		jobNames, err := json.Marshal(list)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.Write(jobNames)
	})

	http.ListenAndServe(":8990", router)
}
