package main

import (
	"github.com/72sevenzy2/http-router/grouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkGrouterDynamicRoutes404(b *testing.B) {
	r := grouter.NewGrouter()

	r.Get("route/path-50/:id", func(w http.ResponseWriter, r *grouter.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/route/invalid-path/72", nil)
	b.ResetTimer()
	b.ReportAllocs()

	for b.Loop() {
		rr := &BenchmarkRR{
			Headers: make(http.Header),
			Code:    http.StatusOK,
		}
		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			b.Fatalf("expected status %d, got %d", http.StatusNotFound, rr.Code)
		}
	}
}

func BenchmarkGolangStdlibRouterDynamicRoutes404(b *testing.B) {
	r := http.NewServeMux()

	r.HandleFunc("GET /route/path-50/{id}", DummyHandler)

	req := httptest.NewRequest(http.MethodGet, "/route/invalid-path/72", nil)

	b.ResetTimer()
	b.ReportAllocs()

	for b.Loop() {
		rr := &BenchmarkRR{
			Headers: make(http.Header),
			Code:    http.StatusOK,
		}
		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			b.Fatalf("expected status %d, got %d", http.StatusNotFound, rr.Code)
		}
	}
}
