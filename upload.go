package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form and create a file to be uploaded
	r.ParseMultipartForm(500 << 20)
	fmt.Println(r.Form)
	filename := r.FormValue("filename")
	uploadedFile, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error uploading file", err)
		return
	}
	defer uploadedFile.Close()

	// Read the first 512 byest to detect the MIME type
	buffer := make([]byte, 512)
	_, err = uploadedFile.Read(buffer)
	if err != nil {
		fmt.Println("Error reading file for MIME type detection:", err)
		return
	}

	// Reset the file reader to the beginning
	uploadedFile.Seek(0, 0)

	// Detect the MIME type
	mimeType := http.DetectContentType(buffer)

	// Determine the file extension based on the MIME type
	var extension string
	switch mimeType {
	case "vide/mp4":
		extension = ".mp4"
	case "video/x-matroska", "video/webm":
		extension = ".mkv"
	default:
		fmt.Println("Unsupported file type:", mimeType)
		return
	}

	//Combine the filename with the extension
	destinationPath := path.Join("uploads", filename+extension)

	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		fmt.Println("Error reaching destination path", err)
		return
	}
	defer destinationFile.Close()

	nBytes, err := io.Copy(destinationFile, uploadedFile)
	if err != nil {
		fmt.Println("Error copying file", err)
		return
	}

	if nBytes == 0 {
		fmt.Println("Warning: Copied file is empty")
	}

	fmt.Println("File uploaded successfully:", filename, "Size:", nBytes, "bytes")

	//Redirect after successful upload
	http.Redirect(w, r, "/", http.StatusSeeOther)

	//TODO have a better upload confirmation screen
	//Successful upload, upload another? go back to the main page
}
