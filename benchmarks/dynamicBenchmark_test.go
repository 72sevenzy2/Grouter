package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/72sevenzy2/http-router/grouter"
)

func BenchmarkGrouterDynamicRoutes(b *testing.B) {
	r := grouter.NewGrouter()

	r.Get("route/path-50/:id", func(w http.ResponseWriter, r *grouter.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/route/path-50/72", nil)
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

// comparing to go's stdlib mux
func BenchmarkGolangStdlibRouterDynamicRoutes(b *testing.B) {
	r := http.NewServeMux()

	r.HandleFunc("GET /route/path-50/{id}", DummyHandler)

	req := httptest.NewRequest(http.MethodGet, "/route/path-50/72", nil)

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

// shared DummyHandler.
func DummyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// shared reusable response recorder (to avoid allocations upon creating a new httptest.NewRecorder() for every request.)
type BenchmarkRR struct {
	Headers http.Header
	Code    int
}

// Header(), WriteHeader(), and Write() implements the http.ResponseWriter interface.
func (w *BenchmarkRR) Header() http.Header {
	return w.Headers
}

func (w *BenchmarkRR) WriteHeader(status int) {
	w.Code = status
}

func (w *BenchmarkRR) Write(p []byte) (int, error) {
	return len(p), nil
}
