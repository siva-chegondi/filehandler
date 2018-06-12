package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/smartsiva/filehandler/store"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// handle submitted form and parse uploaded file
	// FormFile internally will call ParseMultipartForm
	multipartfile, _, err := r.FormFile("file_key")
	if ( err != nil) {
		fmt.Fprintf(w, "\n%s", err)
		return
	}
	defer multipartfile.Close()

	// create minio file instance with uploaded file
	minioFile, err := store.NewMinio("tempname" ,multipartfile)
	if err != nil {
		log.Fatal("Error occured ", err)
		return
	}

	// upload file and return status to user
	fmt.Fprintf(w, minioFile.Upload("smartbucket"))
}

func loadFile(w http.ResponseWriter, r *http.Request) {
	minioFile, err := store.NewMinio("tempname", nil)
	if err != nil {
		log.Fatal("Error occurred", err)
		return
	}
	fileData, err := minioFile.Download("smartbucket")
	if err != nil {
		fmt.Println("Error 1 ", err)
		return
	}
	w.Header().Set("Content-Type", "image/jpg")
	w.Header().Set("Content-Length", "35000")
	w.Write(fileData)
}

func main() {
	
	// set all handlers of server
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/load", loadFile)

	// starting server on port 8087
	http.ListenAndServe(":8087", nil)
}
