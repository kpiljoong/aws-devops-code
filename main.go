package main

import (
	"io"
	"log"
	"net/http"
)

const version string = "0.0.2"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Current Version: "+version)
}

func main() {
	log.Printf("Listening on port 8000...")
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8000", nil)
}
