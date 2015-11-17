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

//not thread safe
var list = make([]string, 0)

func init() {
	fmt.Println("initialize block")
}

func visit(path string, f os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
	}

	matched, err := filepath.Match("*.sh", f.Name())

	if err != nil {
		fmt.Println(err)
	}

	if matched {
		list = append(list, path)
	}

	return nil
}

func main() {
	router := httprouter.New()
	router.GET("/jobnames", func(writer http.ResponseWriter, router *http.Request, _ httprouter.Params) {
		flag.Parse()
		dirRoot := flag.Arg(0)
		err := filepath.Walk(dirRoot, visit)
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
