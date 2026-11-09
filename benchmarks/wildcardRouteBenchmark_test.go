package main

import (
	"github.com/72sevenzy2/grouter/grouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkGrouterWildcardRoute(b *testing.B) {
	r := grouter.NewGrouter()

	r.Get("route/path-50/*someFile", func(w http.ResponseWriter, r *grouter.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/route/path-50/users/somepath/posts", nil)
	b.ResetTimer()
	b.ReportAllocs()

	for b.Loop() {
		rr := &BenchmarkRR{
			Headers: make(http.Header),
			Code:    http.StatusOK,
		}
		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			b.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
		}
	}
}

func BenchmarkGolangStdlibRouterWildCardRoute(b *testing.B) {
	r := http.NewServeMux()

	r.HandleFunc("GET /route/path-50/{someFile...}", DummyHandler)

	req := httptest.NewRequest(http.MethodGet, "/route/path-50/users/somepath/posts", nil)

	b.ResetTimer()
	b.ReportAllocs()

	for b.Loop() {
		rr := &BenchmarkRR{
			Headers: make(http.Header),
			Code:    http.StatusOK,
		}
		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			b.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
		}
	}
}
