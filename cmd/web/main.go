package main

import (
	"fmt"
	"net/http"
)

type application struct{}

func main() {
	app := &application{}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/health", app.healthHandler)
	mux.HandleFunc("/api/v1/health1", app.healthHandler)
	mux.HandleFunc("/api/v1/num/", app.numHandler)

	server := &http.Server{
		Addr:    ":9090",
		Handler: mux,
	}

	fmt.Println("Starting up server on :9090")

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting up server: %s\n", err)
	}
}
