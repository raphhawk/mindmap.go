package main

import (
	"fmt"
	"net/http"
)

type JSONPayload struct {
	Data string `json:"data"`
}

type GitContentPayload struct {
	Name         string            `json:"name"`
	Path         string            `json:"path"`
	Sha          string            `json:"sha"`
	Size         int32             `json:"size"`
	Url          string            `json:"url"`
	Html_url     string            `json:"html_url"`
	Git_url      string            `json:"git_url"`
	Download_url string            `json:"download_url"`
	Type         string            `json:"type"`
	Links        map[string]string `json:"_links"`
}

func (app *Config) GetContents(w http.ResponseWriter, r *http.Request) {
	var requestPayload JSONPayload
	_ = app.readJSON(w, r, &requestPayload)
	requestURL := fmt.Sprintf("https://api.github.com/repos/" + requestPayload.Data + "/contents")
	fmt.Println(requestURL)
	baseDirs, _ := app.gitRequest(requestURL)
	for _, i := range baseDirs {
		fmt.Println(i.Name, i.Type)
	}
	files := []GitContentPayload{}
	app.gitRGrab(baseDirs, files, requestURL)
	fmt.Println(files)
}
