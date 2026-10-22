package routegroupbenchmarks

import (
	"github.com/72sevenzy2/http-router/grouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkInlineRouteGrouping(b *testing.B) {
	r := grouter.NewGrouter()

	api := r.Group("/test1", "child1", "child2")
	api.Get("/child3", func(w http.ResponseWriter, r *grouter.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/test1/child1/child2/child3", nil)

	b.ResetTimer()
	b.ReportAllocs()

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
