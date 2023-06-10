package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"slash/router"
)

const HTTP_SERVER_PORT = ":8080"

func main() {
	err := showASCIILogo("slash-ascii.txt")
	if err != nil {
		log.Fatalf("%s", err)
	}

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

func showASCIILogo(filename string) error {
	readFile, err := os.Open(filename)
	if err != nil {
		return err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	readFile.Close()
	return nil
}
