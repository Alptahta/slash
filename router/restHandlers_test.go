package router

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_createUserHandler(t *testing.T) {
	t.Run("Should return successful response with HTTP status:201", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/users", nil)
		w := httptest.NewRecorder()
		uh := NewUserHandler()
		uh.ServeHTTP(w, req)
		res := w.Result()
		defer res.Body.Close()
		data, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("expected error to be nil got %v", err)
		}
		if string(data) != "TODO create user" {
			t.Fatalf("expected TODO create user got %v", string(data))
		}
		if res.StatusCode != http.StatusCreated {
			t.Fatalf("expected HTTP Status Code 200 got %v", res.StatusCode)
		}
	})
}

func Test_NotFound(t *testing.T) {
	t.Run("Should return HTTP status:404 when POST method not used", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		w := httptest.NewRecorder()
		uh := NewUserHandler()
		uh.ServeHTTP(w, req)
		res := w.Result()
		defer res.Body.Close()
		data, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("expected error to be nil got %v", err)
		}
		if string(data) != "404 page not found\n" {
			t.Fatalf("expected 404 page not found got %v", string(data))
		}
		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("expected HTTP Status Code 404 got %v", res.StatusCode)
		}
	})
}
