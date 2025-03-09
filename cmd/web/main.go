package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
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

	// Connect to the database
	dsn := os.Getenv("MYSQL_DSN")
	db, err := openDB(dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	// Start the server
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}
	logger.Info("starting server", "addr", serverPort)
	serverErr := http.ListenAndServe(fmt.Sprintf(":%s", serverPort), app.routes())
	logger.Error(serverErr.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
