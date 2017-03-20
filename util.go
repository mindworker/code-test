package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WriteError writes JSON error response with given status code
func WriteError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(B(`{"error":"` + err.Error() + `"}`))
}

// PrintJSON prints out any object in JSON format
func PrintJSON(obj interface{}) error {
	jsBuf, err := json.MarshalIndent(obj, "", "   ")
	if err != nil {
		return err
	}

	fmt.Println(string(jsBuf))
	return nil
}

// B converts a string to byte array, just used as a shorter alias
func B(str string) []byte {
	return []byte(str)
}
