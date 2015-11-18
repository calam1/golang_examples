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

func init() {
	fmt.Println("initialize block")
}

func getJobNames(dir string) []string {
	list := make([]string, 0)
	err := filepath.Walk(dir, func(dir string, f os.FileInfo, err error) error {
		matched, err := filepath.Match("*.sh", f.Name())

		if err != nil {
			fmt.Println(err)
		}

		if matched {
			list = append(list, dir)
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return list
}

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
