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

func getJobNames(dir string, writer http.ResponseWriter) {
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

	jobNames, err := json.Marshal(list)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jobNames)

}

func main() {
	router := httprouter.New()
	router.GET("/jobnames", func(writer http.ResponseWriter, router *http.Request, _ httprouter.Params) {
		flag.Parse()
		dirRoot := flag.Arg(0)

		getJobNames(dirRoot, writer)
	})

	http.ListenAndServe(":8990", router)
}
