package main

import (
	"fmt"
	"net/http"
	"github.com/smartsiva/filehandler/file"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	multipartfile, _, err := r.FormFile("file_key")
	defer multipartfile.Close()
	if ( err != nil) {
		fmt.Fprintf(w, "\n%s", err)
		return
	}

	f, _ := file.New(multipartfile)
	fmt.Fprintf(w, "%s", f.Upload())
}

func loadFile(w http.ResponseWriter, r *http.Request) {
}

func main() {

	// set all handlers of server
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/load", loadFile)

	// starting server on port 8087
	http.ListenAndServe(":8087", nil)
}
