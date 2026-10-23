package rateLimiterBenchmarks

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/72sevenzy2/http-router/grouter"
)

func BenchmarkRateLimit_ReqsExceeded(b *testing.B) {
	r := grouter.NewGrouter()
	lim := grouter.NewLimiter(100, 1)
	r.Use(lim.RateLimiter())

	r.Get("test1", func(w http.ResponseWriter, r *grouter.Request) {
		w.WriteHeader(http.StatusOK)
	})
	req := httptest.NewRequest(http.MethodGet, "/test1", nil)
	req.RemoteAddr = "127.0.0.1:9000"
	rr := &BenchmarkRR{
		Headers: make(http.Header),
		Code:    http.StatusOK,
	}

	// draining limit early on
	for range 120 {
		rr.Reset()
		r.ServeHTTP(rr, req)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		rr.Reset()
		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusTooEarly {
			b.Fatalf("got response %d, whilst expected %d", rr.Code, http.StatusTooEarly)
		}
	}
}
