package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/72sevenzy2/http-router/grouter"
)

func TestCancelFunc(t *testing.T) {
	b := grouter.NewGrouter()

	mw, cancel := grouter.Canceller()
	b.Use(mw)

	b.Get("/foo", func(w http.ResponseWriter, r *grouter.Request) {
		select {
		case <-r.Context().Done():
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusFailedDependency)
		}
	})

	req := httptest.NewRequest(http.MethodGet, "/foo", nil)
	rr := httptest.NewRecorder()

	cancel()
	b.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("failed: status %d:", rr.Code)
	}

	t.Log("successful")
}

// auth testing
func TestBasicAuth(t *testing.T) {
	b := grouter.NewGrouter()

	// apply auth middleware
	b.Use(grouter.BasicAuth("user1", "pass1"))

	// b.Handle(http.MethodGet, "/foo1", func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)
	// })

	b.Get("/foo1", func(w http.ResponseWriter, r *grouter.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/foo1", nil)
	req.SetBasicAuth("user1", "pass1")

	rr := httptest.NewRecorder()

	b.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("failed with status %d", rr.Code)
	}
	t.Log("successful.")
}

func TestBearerAuth(t *testing.T) {
	b := grouter.NewGrouter()

	b.Use(grouter.BearerAuth("bearerauth123"))

	b.Get("/foo2", func(w http.ResponseWriter, r *grouter.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/foo2", nil)
	// set auth header and key for BearerAuth()

	req.Header.Set("Authorization", "bearerauth123")

	rr := httptest.NewRecorder()

	b.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("failed auth with status %d", rr.Code)
	}
	t.Log("successful.")
}

// test func to check if logger mw calls next middleware:
func TestLoggerNext(t *testing.T) {
	called := false

	next := func(w http.ResponseWriter, r *grouter.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	}

	handler := grouter.Logger(1024)(next)
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("test"))
	rr := httptest.NewRecorder()
	routerReq := &grouter.Request{Request: req}

	handler(rr, routerReq)

	if !called {
		t.Log("logger did not call next().")
	}

	if rr.Code != http.StatusOK {
		t.Fatalf("failed with status %d", rr.Code)
	}

}

// test to make sure logger preserves data (body)
func TestLoggerBody(t *testing.T) {
	next := func(w http.ResponseWriter, r *grouter.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}

		if string(body) != "testC" {
			t.Fatalf("expected body %q, received %q", "testC", string(body))
		}
	}

	handler := grouter.Logger(1024)(next)

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("testC"))
	rr := httptest.NewRecorder()
	routerReq := &grouter.Request{Request: req}

	handler(rr, routerReq)
}

// rate limiting test
func TestRateLimiter(t *testing.T) {
	b := grouter.NewGrouter()

	lim := grouter.NewLimiter(100, 1) // 100 requests cap, 1 token refill per second

	b.Use(lim.RateLimiter())

	var recs int
	var mu sync.Mutex

	b.Get("/rateLimTest", func(w http.ResponseWriter, r *grouter.Request) {
		mu.Lock()
		recs++
		mu.Unlock()

		w.WriteHeader(http.StatusOK)
	})

	var wg sync.WaitGroup
	for range 100 { // 100 concurrent reqs
		wg.Add(1)

		go func() {
			req := httptest.NewRequest(http.MethodGet, "/rateLimTest", nil)
			rr := httptest.NewRecorder()

			b.ServeHTTP(rr, req)
			wg.Done()
		}()
	}
	wg.Wait()

	t.Log("requests received: ", recs)

}
