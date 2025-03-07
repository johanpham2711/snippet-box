package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

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

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display the snippet with ID %d...", id)
	w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func main() {
	loadEnv()

	http.HandleFunc("/{$}", home)
	http.HandleFunc("/healthz", healthCheck)
	http.HandleFunc("/snippet/view/{id}", snippetView)
	http.HandleFunc("/snippet/create", snippetCreate)

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	log.Println("The application is starting on port", serverPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil)
	log.Fatal(err)
}
