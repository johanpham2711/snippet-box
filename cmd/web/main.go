package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	// Load the environment variables
	loadEnv()

	// Create a new logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Create a new application instance
	app := &application{
		logger: logger,
	}

	// Serve static files
	fileServer := http.FileServer(http.Dir("ui/static/"))
	http.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Start the server
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}
	logger.Info("starting server", "addr", serverPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", serverPort), app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
