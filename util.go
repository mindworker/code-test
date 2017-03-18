package main

import "net/http"

func WriteError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(B(`{"error":"` + err.Error() + `"}`))
}

func B(str string) []byte {
	return []byte(str)
}
