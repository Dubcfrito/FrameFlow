package main

import (
	"fmt"
	"net/http"
)

func main() {
	//Serve the main HTML webpage
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	//Serve the upload webpage
	http.HandleFunc("/upload", UploadHandler)
	http.HandleFunc("/upload-page", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "upload.html")
	})

	//Serve the videos page
	http.HandleFunc("/videos", VideosHandler)
	http.HandleFunc("/videos-page", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "videos.html")
	})

	//Serve CSS files and other static assets
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
