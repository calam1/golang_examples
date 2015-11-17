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

func getJobNames(dir string, c chan []string) {
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

	c <- list
}

func main() {
	router := httprouter.New()
	router.GET("/jobnames", func(writer http.ResponseWriter, router *http.Request, _ httprouter.Params) {
		flag.Parse()
		dirRoot := flag.Arg(0)

		var jobNameChannel chan []string = make(chan []string)
		go getJobNames(dirRoot, jobNameChannel)
		list := <-jobNameChannel
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
