package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func VideosHandler(w http.ResponseWriter, r *http.Request) {
	//Read file names from uploads directory
	files, err := ioutil.ReadDir("uploads")
	if err != nil {
		fmt.Println("Error reading file names", err)
	}

	//Extract file names
	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	tmpl, err := template.ParseFiles("videos.html")
	if err != nil {
		fmt.Println("Error creating template of files", err)
		return
	}

	err = tmpl.Execute(w, fileNames)
	if err != nil {
		fmt.Println("Error doing shit", err)
		return
	}
}
