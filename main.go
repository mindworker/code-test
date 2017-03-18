package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Port defines the port on which the http server will bind to. Note
// that this value gets overwritten by the RAV_PORT environment
var Port = "8080"

// UseCORS tells the Middleware to use CORS. This is for example
// useful if we want to test the API via Postman or curl
var UseCORS = false

func init() {
	port := os.Getenv("RAV_PORT")
	if port != "" {
		Port = port
	}

	if os.Getenv("RAV_USE_CORS") == "true" {
		UseCORS = true
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/actions", handleActions)
	mux.Handle("/", http.FileServer(http.Dir("public")))

	log.Println("Listening on port", Port)
	http.ListenAndServe(":"+Port, &Middleware{mux})
}

func handleActions(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	js := make(map[string]interface{})
	err := decoder.Decode(&js)
	if err != nil {
		WriteError(w, 400, err)
		return
	}

	log.Println(js)
}
