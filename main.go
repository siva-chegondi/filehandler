package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/smartsiva/filehandler/store"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	multipartfile, _, err := r.FormFile("file_key")
	if ( err != nil) {
		fmt.Fprintf(w, "\n%s", err)
		return
	}
	defer multipartfile.Close()

	_, err = store.NewMinio("smarttest", multipartfile)
	if err != nil {
		log.Fatal("Error occured ", err)
		return
	}
	// minioFile.Upload()
}

func loadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Code to load your file")
}

func main() {
	
	// set all handlers of server
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/load", loadFile)

	// starting server on port 8087
	http.ListenAndServe(":8087", nil)
}
