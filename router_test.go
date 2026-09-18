package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/72sevenzy2/http-router/grouter"
)

func TestRouter(t *testing.T) {
	b := grouter.NewGrouter()

	b.Handle(http.MethodGet, "/test/", func(w http.ResponseWriter, r *grouter.Request) {
		w.WriteHeader(http.StatusOK)
	})
	rec := httptest.NewRecorder()
	rr := httptest.NewRequest(http.MethodGet, "/test", nil)

	b.ServeHTTP(rec, rr)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}
