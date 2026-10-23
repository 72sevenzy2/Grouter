package rateLimiterBenchmarks

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/72sevenzy2/http-router/grouter"
)

func BenchmarkRateLimit_allowedReqs(b *testing.B) {
	r := grouter.NewGrouter()
	// massive maximum request limit for testing purposes.
	lim := grouter.NewLimiter(1_000_000_000, 1)
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

	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		rr.Reset()
		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			b.Fatalf("got response %d, whilst expected %d", rr.Code, http.StatusOK)
		}
	}
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

func (w *BenchmarkRR) Reset() {
	w.Code = 0
}
