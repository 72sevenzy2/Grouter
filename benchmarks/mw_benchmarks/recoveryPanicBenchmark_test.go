package mwbenchmarks

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/72sevenzy2/http-router/grouter"
)

func BenchmarkRecoveryMiddlewarePanic(b *testing.B) {
	slog.SetDefault(slog.New(slog.NewTextHandler(io.Discard, nil)))
	r := grouter.NewGrouter()
	// measures recovery middleware overhead whilst no panic occurs.
	r.Use(grouter.Recoverer())

	r.Get("test1", func(w http.ResponseWriter, r *grouter.Request) {
		panic("very bad server crash.")
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

		if rr.Code != http.StatusInternalServerError {
			b.Fatalf("got response %d, whilst expected %d", rr.Code, http.StatusInternalServerError)
		}
	}
}
