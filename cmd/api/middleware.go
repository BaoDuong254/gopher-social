package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func (app *application) BasicAuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// read the auth header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				app.unauthorizedBasicErrorResponse(w, r, fmt.Errorf("missing authorization header"))
				return
			}

			// parse it
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Basic" {
				app.unauthorizedBasicErrorResponse(w, r, fmt.Errorf("invalid authorization header"))
				return
			}

			// decode the base64 encoded credentials
			decoded, err := base64.StdEncoding.DecodeString(parts[1])
			if err != nil {
				app.unauthorizedBasicErrorResponse(w, r, err)
				return
			}
			username := app.config.auth.basic.user
			password := app.config.auth.basic.pass
			creds := strings.Split(string(decoded), ":")
			if len(creds) != 2 || creds[0] != username || creds[1] != password {
				app.unauthorizedBasicErrorResponse(w, r, fmt.Errorf("invalid credentials"))
				return
			}

			// call the next handler in the chain
			next.ServeHTTP(w, r)
		})
	}
}
