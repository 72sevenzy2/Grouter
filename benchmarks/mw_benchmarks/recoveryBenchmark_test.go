package mwbenchmarks

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/72sevenzy2/http-router/grouter"
)

func BenchmarkRecoveryMiddleware(b *testing.B) {
	r := grouter.NewGrouter()
	// measures recovery middleware overhead whilst no panic occurs.
	r.Use(grouter.Recoverer())

	r.Get("test1", func(w http.ResponseWriter, r *grouter.Request) {
		w.WriteHeader(http.StatusOK)
	})

	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		rr := &BenchmarkRR{
			Headers: make(http.Header),
			Code:    http.StatusOK,
		}

		req := httptest.NewRequest(http.MethodGet, "/test1", nil)

		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			b.Fatalf("got response %d, whilst expected %d", rr.Code, http.StatusOK)
		}
	}
}
