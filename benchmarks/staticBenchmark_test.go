package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/72sevenzy2/http-router/grouter"
)

func BenchmarkGrouterStaticRoutes(t *testing.B) {
	t.ReportAllocs()
	r := grouter.NewGrouter()

	r.Get("test/path/hi", func(w http.ResponseWriter, r *grouter.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/test/path/hi", nil)
	rr := httptest.NewRecorder()

	t.ResetTimer()

	for t.Loop() {
		r.ServeHTTP(rr, req)
	}
}

func BenchmarkGolangStdlibStaticRoutes(t *testing.B) {
	t.ReportAllocs()
	r := http.NewServeMux()

	r.HandleFunc("test/path/hi", DummyHandler)
	req := httptest.NewRequest(http.MethodGet, "/test/path/hi", nil)
	rr := httptest.NewRecorder()

	t.ResetTimer()
	for t.Loop() {
		r.ServeHTTP(rr, req)
	}
}
