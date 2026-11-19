package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type stubPinger struct {
	configured bool
	err        error
}

func (s stubPinger) Configured() bool {
	return s.configured
}

func (s stubPinger) Ping(context.Context) error {
	return s.err
}

func TestRouterEndpoints(t *testing.T) {
	t.Parallel()

	t.Run("GET /healthz returns 200 ok with no-store", func(t *testing.T) {
		t.Parallel()
		router := NewRouter("development", []string{"http://localhost:3000"}, stubPinger{})
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
		}
		if got := rec.Header().Get("Cache-Control"); got != "no-store" {
			t.Fatalf("Cache-Control = %q, want no-store", got)
		}
		if rec.Header().Get("X-Request-Id") == "" {
			t.Fatal("expected X-Request-Id header to be set")
		}
	})

	t.Run("GET /readyz returns 200 when database reachable", func(t *testing.T) {
		t.Parallel()
		router := NewRouter("production", []string{"http://localhost:3000"}, stubPinger{configured: true})
		req := httptest.NewRequest(http.MethodGet, "/readyz", nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
		}
	})

	t.Run("GET /readyz fails closed with RFC 9457 when database down", func(t *testing.T) {
		t.Parallel()
		router := NewRouter("production", []string{"http://localhost:3000"}, stubPinger{
			configured: true,
			err:        errors.New("dial tcp timeout"),
		})
		req := httptest.NewRequest(http.MethodGet, "/readyz", nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusServiceUnavailable {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusServiceUnavailable)
		}
		var problem ProblemDetails
		if err := json.NewDecoder(rec.Body).Decode(&problem); err != nil {
			t.Fatalf("decode problem json: %v", err)
		}
		if problem.Status != http.StatusServiceUnavailable {
			t.Fatalf("problem.status = %d, want 503", problem.Status)
		}
	})

	t.Run("GET /readyz in production fails closed without database", func(t *testing.T) {
		t.Parallel()
		router := NewRouter("production", []string{"http://localhost:3000"}, stubPinger{configured: false})
		req := httptest.NewRequest(http.MethodGet, "/readyz", nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusServiceUnavailable {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusServiceUnavailable)
		}
	})

	t.Run("GET /api/v1/public/ping returns Hello SINERGI", func(t *testing.T) {
		t.Parallel()
		router := NewRouter("development", []string{"http://localhost:3000"}, stubPinger{})
		req := httptest.NewRequest(http.MethodGet, "/api/v1/public/ping", nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
		}
		var body map[string]string
		if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		if body["message"] != "Hello SINERGI" {
			t.Fatalf("message = %q, want Hello SINERGI", body["message"])
		}
	})

	t.Run("GET /admin returns uniform 404 RFC 9457 (anti-enumeration)", func(t *testing.T) {
		t.Parallel()
		router := NewRouter("production", []string{"http://localhost:3000"}, stubPinger{})
		req := httptest.NewRequest(http.MethodGet, "/admin", nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusNotFound {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusNotFound)
		}
	})
}
