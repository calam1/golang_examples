package main

type FileName struct {
	Name           string `json:"name"`
	Path           string `json:"path"`
	FileContentsUrl string `json:"filecontentsurl"`
}

type Files struct {
	FileNames []FileName `json:"filenames"`
}
