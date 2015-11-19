package main

type JobName struct {
	Name           string `json:"name"`
	Path           string `json:"path"`
	JobContentsUrl string `json:"jobcontentsurl"`
	JobExecuteUrl  string `json:"jobExecuteurl"`
}

type Jobs struct {
	JobNames []JobName `json:"jobnames"`
}
