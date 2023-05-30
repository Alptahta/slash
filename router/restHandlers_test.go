package router

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_createUserHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/users", nil)
	w := httptest.NewRecorder()
	uh := NewUserHandler()
	uh.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "TODO create user" {
		t.Errorf("expected TODO create user got %v", string(data))
	}
}

func Test_NotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()
	uh := NewUserHandler()
	uh.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	fmt.Println("a" + string(data) + "a")
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "404 page not found\n" {
		t.Errorf("expected 404 page not found got %v", string(data))
	}
}
