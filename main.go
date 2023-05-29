package main

import (
	"log"
	"net/http"
)

const HTTP_SERVER_PORT = "8080"

func main() {
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    HTTP_SERVER_PORT,
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
	log.Printf("Starting HTTP server at port %s", HTTP_SERVER_PORT)
	mux.HandleFunc("/users", createUser)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	resp := "TODO create user"
	w.Write([]byte(resp))
}
