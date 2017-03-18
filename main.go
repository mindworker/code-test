package main

import (
	"log"
	"net/http"
	"os"
)

// Port defines the port on which the http server will bind to. Note
// that this value gets overwritten by the RAV_PORT environment
var Port = "8080"

func init() {
	port := os.Getenv("RAV_PORT")
	if port != "" {
		Port = port
	}
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	log.Println("Listening on port", Port)
	http.ListenAndServe(":"+Port, nil)
}
