package main

import (
	"log"
	"net/http"
)

// Middleware is a wrapper for router, which allows to add
// interceptors to every request, e.g. to add CORS headers
type Middleware struct {
	mux *http.ServeMux
}

// ServeHTTP intercepts requests before passing it on to the actual
// handler function
func (m *Middleware) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fn := m.mux.ServeHTTP

	// Note that UseCORS is defined in main.go
	if UseCORS {
		log.Println("Use CORS")
		fn = CORSWrapper(fn)
	}

	// Can add additional middleware interceptors here
	fn = RestrictionWrapper(fn)

	fn(rw, req)
}

// CORSWrapper generically wraps any http handler and adds CORS
// headers
func CORSWrapper(fn http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		fn(rw, req)
	}
}

// RestrictionWrapper is used to block any other methods but GET and
// POST
func RestrictionWrapper(fn http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET")
		fn(rw, req)
	}
}
