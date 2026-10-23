package rateLimiterBenchmarks

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/72sevenzy2/http-router/grouter"
)

func BenchmarkRateLimit_multiClients(b *testing.B) {
	r := grouter.NewGrouter()
	lim := grouter.NewLimiter(100, 1)
	r.Use(lim.RateLimiter())
	r.Get("test1", func(w http.ResponseWriter, r *grouter.Request) {
		w.WriteHeader(http.StatusOK)
	})

	reqs := make([]*http.Request, 1000)
	for i := range 1000 {
		reqs[i] = httptest.NewRequest(http.MethodGet, "/test1", nil)
		reqs[i].RemoteAddr = fmt.Sprintf("requestno %d", i)
	}

	rr := &BenchmarkRR{
		Headers: make(http.Header),
		Code:    http.StatusOK,
	}
	i := 0

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		rr.Reset()
		r.ServeHTTP(rr, reqs[i])
		i++

		if i == len(reqs) {
			i = 0
		}
	}
}
