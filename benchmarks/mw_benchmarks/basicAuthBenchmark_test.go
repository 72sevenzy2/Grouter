package mwbenchmarks

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/72sevenzy2/http-router/grouter"
)

func BenchmarkBasicAuthMiddleware(b *testing.B) {
	r := grouter.NewGrouter()
	r.Use(grouter.BasicAuth("user1", "pass1"))

	r.Get("test1", func(w http.ResponseWriter, r *grouter.Request) {
		w.WriteHeader(http.StatusOK)
	})
	req := httptest.NewRequest(http.MethodGet, "/test1", nil)
	req.SetBasicAuth("user1", "pass1")

	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		rr := &BenchmarkRR{
			Headers: make(http.Header),
			Code:    http.StatusOK,
		}
		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			b.Fatalf("got response %d, whilst expected %d", rr.Code, http.StatusOK)
		}
	}
}
