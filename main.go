package main

import (
	"fmt"
	"log"
	"net/http"

	"embed"
	"slash/router"
)

const HTTP_SERVER_PORT = ":8080"

//go:embed slash-ascii.txt
var f embed.FS

func main() {
	err := showASCIILogo("slash-ascii.txt")
	if err != nil {
		log.Fatalf("%s", err)
	}

	mux := http.NewServeMux()

	uh := router.NewUserHandler()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/users", uh.ServeHTTP)

	// TODO Improve Server fields
	srv := &http.Server{
		Addr:    HTTP_SERVER_PORT,
		Handler: mux,
		// ReadTimeout: time.Duration(2),
	}

	log.Printf("Starting HTTP server at port %s", HTTP_SERVER_PORT)
	log.Fatal(srv.ListenAndServe())
}

func showASCIILogo(filename string) error {
	data, err := f.ReadFile(filename)
	if err != nil {
		return err
	}
	_, err = fmt.Println(string(data))
	if err != nil {
		return err
	}
	return nil
}
