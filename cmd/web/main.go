package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load the environment variables
	loadEnv()

	// Serve static files
	fileServer := http.FileServer(http.Dir("ui/static/"))
	http.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Register the route handlers
	http.HandleFunc("GET /{$}", home)
	http.HandleFunc("GET /healthz", healthCheck)
	http.HandleFunc("GET /snippet/view/{id}", snippetView)
	http.HandleFunc("POST /snippet/create", snippetCreate)

	// Start the server
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}
	log.Println("The application is starting on port", serverPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil)
	log.Fatal(err)
}
