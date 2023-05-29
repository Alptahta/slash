package router

import (
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
