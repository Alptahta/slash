package router

import "net/http"

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
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
	w.Write([]byte(resp))
}
