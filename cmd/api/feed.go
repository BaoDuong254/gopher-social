package main

import (
	"errors"
	"net/http"

	"github.com/baoduong254/gopher-social/internal/store"
)

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	fq := store.PaginatedFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}
	fq, err := fq.Parse(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	if err := Validate.Struct(fq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	user := app.getUserFromContext(r)
	if user == nil {
		app.unauthorizedResponse(w, r, errors.New("missing authenticated user in request context"))
		return
	}

	feed, err := app.store.Posts.GetUserFeed(r.Context(), user.ID, fq)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	if err := app.jsonResponse(w, http.StatusOK, feed); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
