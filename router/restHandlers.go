package router

import (
	"net/http"
	"slash/service"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (uh *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		uh.createUserHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (uh *UserHandler) createUserHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement me
	resp := "TODO create user"
	w.WriteHeader(http.StatusCreated)
	_, err := w.Write([]byte(resp))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
