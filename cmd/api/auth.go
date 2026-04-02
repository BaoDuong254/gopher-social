package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"net/http"

	"github.com/baoduong254/gopher-social/internal/mailer"
	"github.com/baoduong254/gopher-social/internal/store"
	"github.com/google/uuid"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}

type UserWithToken struct {
	*store.User
	Token string `json:"token"`
}

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	user := &store.User{
		Username: payload.Username,
		Email:    payload.Email,
	}

	// Hash the password and set the PasswordHash field on the user struct.
	if err := user.Password.Set(payload.Password); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	plainToken := uuid.New().String()
	hash := sha256.Sum256([]byte(plainToken))

	err := app.store.Users.CreateAndInvite(r.Context(), user, hash[:], app.config.mail.exp)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrDuplicateUsername):
			app.badRequestResponse(w, r, err)
		case errors.Is(err, store.ErrDuplicateEmail):
			app.badRequestResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}
	userWithToken := &UserWithToken{
		User:  user,
		Token: plainToken,
	}
	activationURL := fmt.Sprintf("%s/confirm/%s", app.config.frontendURL, plainToken)
	isProdEnv := app.config.env == "production"
	vars := struct {
		Username      string
		ActivationURL string
	}{
		Username:      user.Username,
		ActivationURL: activationURL,
	}

	err = app.mailer.Send(mailer.UserWelcomeEmailTemplate, user.Username, user.Email, vars, !isProdEnv)
	if err != nil {
		if err := app.store.Users.Delete(r.Context(), user.ID); err != nil {
			app.logger.Errorf("failed to delete user after email sending failure: %v", err)
		}
		app.internalServerError(w, r, err)
		return
	}
	if err := app.jsonResponse(w, http.StatusCreated, userWithToken); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
