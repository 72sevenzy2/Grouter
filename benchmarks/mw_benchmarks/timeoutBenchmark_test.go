package mwbenchmarks

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/72sevenzy2/http-router/grouter"
)

func BenchmarkTimeoutMiddleware(b *testing.B) {
	r := grouter.NewGrouter()
	r.Use(grouter.Timeout(time.Second * 2)) // abort request after 2 seconds

	r.Get("test1", func(w http.ResponseWriter, r *grouter.Request) {
		select {
		case <-time.After(10 * time.Second):
			w.WriteHeader(http.StatusOK)
		case <-r.Context().Done():
			w.WriteHeader(http.StatusBadRequest)
			return
		}
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

		if rr.Code != http.StatusBadRequest {
			b.Fatalf("got response %d, whilst expected %d", rr.Code, http.StatusBadRequest)
		}
	}
}
