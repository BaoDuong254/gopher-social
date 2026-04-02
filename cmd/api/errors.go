package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err.Error())
	errs := writeJSONError(w, http.StatusInternalServerError, "internal server error")
	if errs != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error: %s path: %s error: %s", r.Method, r.URL.Path, err.Error())
	errs := writeJSONError(w, http.StatusBadRequest, err.Error())
	if errs != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("not found error: %s path: %s error: %s", r.Method, r.URL.Path, err.Error())
	errs := writeJSONError(w, http.StatusNotFound, err.Error())
	if errs != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("conflict error: %s path: %s error: %s", r.Method, r.URL.Path, err.Error())
	errs := writeJSONError(w, http.StatusConflict, err.Error())
	if errs != nil {
		http.Error(w, err.Error(), http.StatusConflict)
	}
}

func (app *application) unauthorizedResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("unauthorized error: %s path: %s error: %s", r.Method, r.URL.Path, err.Error())
	errs := writeJSONError(w, http.StatusUnauthorized, err.Error())
	if errs != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
}

func (app *application) unauthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("unauthorized error: %s path: %s error: %s", r.Method, r.URL.Path, err.Error())
	w.Header().Set("WWW-Authenticate", `Basic realm="Restricted", charset="UTF-8"`)
	errs := writeJSONError(w, http.StatusUnauthorized, "invalid credentials")
	if errs != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
	}
}
