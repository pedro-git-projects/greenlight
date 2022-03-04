package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// hard-coded now, but will be generated at build time in the future
const version = "1.0.0"

// stores our configuration setting for the application
// only storing the port to listen to and the environment type atm
type config struct {
	port int
	env  string
}

// holds the dependencies for HTTP handlers, helpers and middleware
// contains the config and a logger, but shall grow with the program
type application struct {
	config config
	logger *log.Logger
}

func main() {
	// instantiating our configuration
	var cfg config

	// if no flags are passed the default port is 8080 and the default environment is development
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	// Initializing a new logger that prefixes its output with the date and time
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// instantiating our application with the defined cfg and logger
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Declaring an http server with sensible timeout setting which will listen on the port provided in the config struct
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Starting the HTTP server
	logger.Printf("Starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
