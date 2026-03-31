package main

import "net/http"

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "available",
		"env":     app.config.env,
		"version": version,
	}
	if err := writeJSON(w, http.StatusOK, data); err != nil {
		err := writeJSONError(w, http.StatusInternalServerError, "failed to write JSON response")
		if err != nil {
			http.Error(w, "failed to write JSON response", http.StatusInternalServerError)
		}
	}
}
