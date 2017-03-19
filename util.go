package main

import "net/http"

// WriteError writes JSON error response with given status code
func WriteError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(B(`{"error":"` + err.Error() + `"}`))
}

// B converts a string to byte array, just used as a shorter alias
func B(str string) []byte {
	return []byte(str)
}
