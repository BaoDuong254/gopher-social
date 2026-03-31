package main

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return err
	}
	return nil
}

// func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
// 	maxBytes := 1_048_576                                    // 1MB
// 	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes)) // Limit the size of the request body to prevent malicious requests from consuming too much memory.
// 	decoder := json.NewDecoder(r.Body)
// 	decoder.DisallowUnknownFields() // Disallow unknown fields in the JSON request body.
// 	err := decoder.Decode(data)
// 	if err != nil {
// 		http.Error(w, "failed to decode JSON request body", http.StatusBadRequest)
// 		return err
// 	}
// 	return nil
// }

func writeJSONError(w http.ResponseWriter, status int, message string) error {
	type envelope struct {
		Error string `json:"error"`
	}
	err := writeJSON(w, status, &envelope{Error: message})
	if err != nil {
		return err
	}
	return nil
}
