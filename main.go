package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippet Box"))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The application is healthy!"))
}

func main() {
	loadEnv()
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/healthz", healthCheck)

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	log.Println("The application is starting on port", serverPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", serverPort), mux)
	log.Fatal(err)
}
