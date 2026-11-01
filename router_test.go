package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/72sevenzy2/http-router/grouter"
)

func TestStaticRoutes(t *testing.T) {
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

func TestStaticRoutes404(t *testing.T) {
	b := grouter.NewGrouter()

	b.Handle(http.MethodGet, "/test/", func(w http.ResponseWriter, r *grouter.Request) {
		w.WriteHeader(http.StatusOK)
	})
	rec := httptest.NewRecorder()
	rr := httptest.NewRequest(http.MethodGet, "/test/a", nil)

	b.ServeHTTP(rec, rr)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}
}

func TestDynamicRoutes(t *testing.T) {
	b := grouter.NewGrouter()

	b.Handle(http.MethodGet, "/test/:id", func(w http.ResponseWriter, r *grouter.Request) {
		for k, v := range r.Params {
			t.Log(k, v)
		}
		w.WriteHeader(http.StatusOK)
	})
	rec := httptest.NewRecorder()
	rr := httptest.NewRequest(http.MethodGet, "/test/10", nil)

	b.ServeHTTP(rec, rr)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestDynamicRoutes404(t *testing.T) {
	b := grouter.NewGrouter()

	b.Handle(http.MethodGet, "/test/:id", func(w http.ResponseWriter, r *grouter.Request) {
		for k, v := range r.Params {
			t.Log(k, v)
		}
		w.WriteHeader(http.StatusOK)
	})
	rec := httptest.NewRecorder()
	rr := httptest.NewRequest(http.MethodGet, "/test1/10", nil)

	b.ServeHTTP(rec, rr)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}
}

func TestWildcardRoute(t *testing.T) {
	b := grouter.NewGrouter()

	b.Handle(http.MethodGet, "/test/*filepath", func(w http.ResponseWriter, r *grouter.Request) {
		w.WriteHeader(http.StatusOK)
	})
	rec := httptest.NewRecorder()
	rr := httptest.NewRequest(http.MethodGet, "/test/someFile", nil)

	b.ServeHTTP(rec, rr)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

}
