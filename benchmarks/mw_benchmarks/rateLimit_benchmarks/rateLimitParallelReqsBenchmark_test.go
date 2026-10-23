package rateLimiterBenchmarks

import (
	"github.com/72sevenzy2/http-router/grouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkRateLimit_parallelReqs(b *testing.B) {
	r := grouter.NewGrouter()
	lim := grouter.NewLimiter(100, 1)
	r.Use(lim.RateLimiter())
	r.Get("test1", func(w http.ResponseWriter, r *grouter.Request) {
		w.WriteHeader(http.StatusOK)
	})

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(p *testing.PB) {
		req := httptest.NewRequest(http.MethodGet, "/test1", nil)
		req.RemoteAddr = "127.0.0.1:9000"
		rr := &BenchmarkRR{
			Headers: make(http.Header),
			Code:    http.StatusOK,
		}

		for p.Next() {
			rr.Reset()
			r.ServeHTTP(rr, req)
		}
	})
}
