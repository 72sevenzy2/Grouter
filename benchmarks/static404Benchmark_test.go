package main

import (
	"github.com/72sevenzy2/http-router/grouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkGrouterStaticRoutes404(b *testing.B) {
	r := grouter.NewGrouter()

	r.Get("test/path/hi", func(w http.ResponseWriter, r *grouter.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/test/invalid-path/hi", nil)

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

func BenchmarkGolangStdlibStaticRoutes404(b *testing.B) {
	r := http.NewServeMux()

	r.HandleFunc("GET /test/path/hi", DummyHandler)
	req := httptest.NewRequest(http.MethodGet, "/test/invalid-path/hi", nil)

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
