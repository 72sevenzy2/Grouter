package mwbenchmarks

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/72sevenzy2/http-router/grouter"
)

func BenchmarkLoggerMiddleware(b *testing.B) {
	slog.SetDefault(slog.New(slog.NewTextHandler(io.Discard, nil))) // discard actual request logs.

	r := grouter.NewGrouter()
	r.Use(grouter.Logger(0)) // set as default request body logging size ( which is 0.)

	r.Get("test1", func(w http.ResponseWriter, r *grouter.Request) {
		w.WriteHeader(http.StatusOK)
	})
	req := httptest.NewRequest(http.MethodGet, "/test1", nil)

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
