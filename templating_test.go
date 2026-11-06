package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/72sevenzy2/grouter/grouter"
)

func TestGrouter403Template(t *testing.T) {
	r := grouter.NewGrouter()

	r.Get("test/test1", func(w http.ResponseWriter, r *grouter.Request) {
		grouter.DisplayForbidden(w)
	})

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test/test1", nil)

	r.ServeHTTP(rr, req)

	assertResponseMetadata(rr, http.StatusForbidden)
}

func TestGrouter404Template(t *testing.T) {
	r := grouter.NewGrouter()

	r.Get("test/test1", func(w http.ResponseWriter, r *grouter.Request) {
		grouter.DisplayNotFound(w)
	})

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test/test1", nil)

	r.ServeHTTP(rr, req)

	assertResponseMetadata(rr, http.StatusNotFound)
}

func TestGrouter401Template(t *testing.T) {
	r := grouter.NewGrouter()

	r.Get("test/test1", func(w http.ResponseWriter, r *grouter.Request) {
		grouter.DisplayUnauthorized(w)
	})

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test/test1", nil)

	r.ServeHTTP(rr, req)

	assertResponseMetadata(rr, http.StatusUnauthorized)
}

func TestGrouter500Template(t *testing.T) {
	r := grouter.NewGrouter()

	r.Get("test/test1", func(w http.ResponseWriter, r *grouter.Request) {
		grouter.DisplayInternalServerError(w)
	})

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test/test1", nil)

	r.ServeHTTP(rr, req)

	assertResponseMetadata(rr, http.StatusInternalServerError)
}

func assertResponseMetadata(resp *httptest.ResponseRecorder, statusCode int) error {
	if resp.Result().Header.Get("content-type") != "text/html" {
		return fmt.Errorf("expected content-type %s, got %s", "text/html", resp.Result().Header.Get("content-type"))
	}
	if resp.Result().StatusCode != statusCode {
		return fmt.Errorf("expected status code %d, got %d", statusCode, resp.Result().StatusCode)
	}
	return nil
}
