package main

import (
	"log"
	"net/http"

	"slash/router"
)

const HTTP_SERVER_PORT = ":8080"

func main() {
	mux := http.NewServeMux()

	uh := router.NewUserHandler()

	mux.HandleFunc("/users", uh.ServeHTTP)

	srv := &http.Server{
		Addr:    HTTP_SERVER_PORT,
		Handler: mux,
	}

	log.Printf("Starting HTTP server at port %s", HTTP_SERVER_PORT)
	log.Fatal(srv.ListenAndServe())
}
