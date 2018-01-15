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

func FilesInfo(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	//we can take a commmand line argument if one is passed
	flag.Parse()
	dirRoot := flag.Arg(0)

	//if url is passed this will override the passed argument at start up
	directory := request.FormValue("directory")
	prettyPrint := request.FormValue("prettyprint")

	if directory != "" {
		dirRoot = directory
	}

  //fmt.Printf("Req: %s %s\n", request.Host, request.URL.Path) 

	fNames := getFNames(dirRoot, request.Host)
	fs := Files{fNames}

	var fsJson []byte
	var err error
	if prettyPrint == "true" {
		fsJson, err = json.MarshalIndent(fs, "", "    ")
	} else {
		fsJson, err = json.Marshal(fs)
	}

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(fsJson)
}

//use of closure and anonymous function
func getFNames(directory string, host string) []FileName {
	fNames := make([]FileName, 0)
	err := filepath.Walk(directory, func(dir string, f os.FileInfo, err error) error {
		matched, err := filepath.Match("*.*", f.Name())

		if err != nil {
			fmt.Println(err)
		}

		if matched {
			dirName, fileName := filepath.Split(dir)
      contentUrl := "http://" + host + "/files/file/" + fileName + "?directory=" + dirName
			fName := FileName{fileName, dirName, contentUrl}
			fNames = append(fNames, fName)
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return fNames
}
