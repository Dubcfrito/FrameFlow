package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data to retrieve the file
	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile("video")
	if err != nil {
		fmt.Println("Error retrieving file:", err)
		return
	}
	defer file.Close()

	// Create a new file in the server to store the uploaded file
	dst, err := os.Create("uploads/video.mp4") // You'll want to create a unique name for each video
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the new file
	if _, err := io.Copy(dst, file); err != nil {
		fmt.Println("Error copying file:", err)
		return
	}

	//Redirect or inform the user that the upload was succesful
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
