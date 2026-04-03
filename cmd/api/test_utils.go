package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/baoduong254/gopher-social/internal/auth"
	"github.com/baoduong254/gopher-social/internal/store"
	"github.com/baoduong254/gopher-social/internal/store/cache"
	"go.uber.org/zap"
)

func newTestApplication(t *testing.T) *application {
	t.Helper()
	logger := zap.NewNop().Sugar()
	mockStore := store.NewMockStorage()
	mockCacheStore := cache.NewMockStorage()
	testAuth := &auth.TestAuthenticator{}
	return &application{
		logger:        logger,
		store:         mockStore,
		cacheStorage:  mockCacheStore,
		authenticator: testAuth,
	}
}

func executeRequest(req *http.Request, mux http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d, got %d", expected, actual)
	}
}
