package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

var list = make([]string, 0)

//prints all values; directories and files
//func visit(path string, f os.FileInfo, err error) error {
//	fmt.Printf("Visited: %s\n", path)
//	return nil
//}

//this is recursive, and filters out .txt files
func visit(path string, f os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
	}

	//filter out what type of files we want to list
	matched, err := filepath.Match("*.txt", f.Name())

	if err != nil {
		fmt.Println(err)
		return err
	}

	if matched {
		list = append(list, path)
		fmt.Println(path)
	}

	return nil
}

//walks the whole directory structure of the passed in root recursively
//func main() {
//	flag.Parse()
//	root := flag.Arg(0)
//	err := filepath.Walk(root, visit)
//	fmt.Printf("filepath.Walk() returned %v\n", err)
//}

//walks the directory and builds json to print out
func main() {
	flag.Parse()
	root := flag.Arg(0)
	err := filepath.Walk(root, visit)
	fmt.Printf("the list size is: %d\n", len(list))
	fmt.Println(list)
	fmt.Printf("filepath.Walk() returned %v\n", err)

	for i := range list {
		fmt.Println(list[i])
		fmt.Println(reflect.TypeOf(list[i]))
	}

	listJson, _ := json.Marshal(list)
	fmt.Println(string(listJson))
}
