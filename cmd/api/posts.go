package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/baoduong254/gopher-social/internal/store"
	"github.com/go-chi/chi/v5"
)

type CreatePostPayload struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload
	if err := readJSON(w, r, &payload); err != nil {
		if err := writeJSONError(w, http.StatusBadRequest, error.Error(err)); err != nil {
			http.Error(w, error.Error(err), http.StatusBadRequest)
		}
		return
	}
	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		UserID:  1,
		Tags:    payload.Tags,
	}
	ctx := r.Context()
	if err := app.store.Posts.Create(ctx, post); err != nil {
		err := writeJSONError(w, http.StatusInternalServerError, err.Error())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		err := writeJSONError(w, http.StatusInternalServerError, err.Error())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		err := writeJSONError(w, http.StatusBadRequest, "invalid post ID")
		if err != nil {
			http.Error(w, "invalid post ID", http.StatusBadRequest)
		}
		return
	}
	ctx := r.Context()
	post, err := app.store.Posts.GetByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			err := writeJSONError(w, http.StatusNotFound, "post not found")
			if err != nil {
				http.Error(w, "post not found", http.StatusNotFound)
			}
		default:
			err := writeJSONError(w, http.StatusInternalServerError, "failed to retrieve post")
			if err != nil {
				http.Error(w, "failed to retrieve post", http.StatusInternalServerError)
			}
		}
		return
	}
	if err := writeJSON(w, http.StatusOK, post); err != nil {
		err := writeJSONError(w, http.StatusInternalServerError, "failed to write JSON response")
		if err != nil {
			http.Error(w, "failed to write JSON response", http.StatusInternalServerError)
		}
		return
	}
}
