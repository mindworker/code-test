package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
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
		log.Println("Using CORS")
		UseCORS = true
	}
}

// Sessions defines our SessionManager
// @TODO might be a good idea to make this a singleton
var Sessions = NewSessionManager(6 * time.Second)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/sessions", PostOnlyWrapper(createSessionHandler))
	mux.HandleFunc("/actions", PostOnlyWrapper(addActionHandler))
	mux.Handle("/", http.FileServer(http.Dir("public")))

	log.Println("Listening on port", Port)
	http.ListenAndServe(":"+Port, &Middleware{mux})
}

func createSessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionID, err := Sessions.CreateSession()
	if err != nil {
		WriteError(w, 500, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(B(`{"sessionId":"` + sessionID + `"}`))
}

func addActionHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	data := make(map[string]interface{})
	err := decoder.Decode(&data)
	if err != nil {
		WriteError(w, 400, err)
		return
	}
}
