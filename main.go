package main

import (
	"fmt"
	"log"
	"net/http"
)

const HTTP_SERVER_PORT = "8080"

func main() {
	log.Printf("Starting HTTP server at port %s", HTTP_SERVER_PORT)
	mux := http.NewServeMux()
	mux.HandleFunc("/users", createUser)

	portString := fmt.Sprintf(":%s", HTTP_SERVER_PORT)
	log.Fatal(http.ListenAndServe(portString, mux))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	resp := "TODO create user"
	w.Write([]byte(resp))
}
