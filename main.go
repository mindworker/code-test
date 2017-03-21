package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
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
var Sessions = NewSessionManager(60 * 60 * time.Second)

func main() {
	mux := http.NewServeMux()

	// API handlers
	mux.HandleFunc("/sessions", PostOnlyWrapper(createSessionHandler))
	mux.HandleFunc("/actions", PostOnlyWrapper(addActionHandler))

	// This is used to serve our frontend
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

// AddActionReq defines commonly used field for all action requests
type AddActionReq struct {
	SessionID  string `json:"sessionId"`
	WebsiteURL string `json:"websiteUrl"`
	EventType  string `json:"eventType"`
}

func addActionHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteError(w, 400, err)
		return
	}
	defer r.Body.Close()

	var data AddActionReq
	err = json.Unmarshal(body, &data)
	if err != nil {
		WriteError(w, 400, err)
		return
	}

	sessionData, err := Sessions.GetSessionData(data.SessionID)
	if err != nil {
		WriteError(w, 400, err)
		return
	}

	// Do nothing if form already submitted
	if sessionData.FormCompletionTime != 0 {
		return
	}

	switch data.EventType {
	case "copyAndPaste":
		err = copyAndPaste(body, sessionData)
	case "resizeWindow":
		err = resizeWindow(body, sessionData)
	case "timeTaken":
		err = timeTaken(body, sessionData)
	default:
		err = errors.New("eventType not recognized")
	}
	if err != nil {
		WriteError(w, 400, err)
		return
	}

	sessionData.WebsiteURL = data.WebsiteURL
	err = Sessions.UpdateSessionData(data.SessionID, sessionData)
	if err != nil {
		WriteError(w, 400, err)
		return
	}

	PrintJSON(sessionData)
	if err != nil {
		WriteError(w, 400, err)
		return
	}

	w.WriteHeader(200)
}
