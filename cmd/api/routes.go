package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// By defining routes as a method of our application we can easily access
// our handlers trough the application struct
func (app *application) routes() *httprouter.Router {
	// initializing a new httprouter instance
	router := httprouter.New()

	// configuring router to use our custom error treatment
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// Register the relevant methods, URL patterns and handler functions for our
	// endpoints using the HandlerFunc() method. Note that http.MethodGet and
	// http.MethodPost are constants which equate to the strings "GET" and "POST"
	// respectively.
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	// :<parameter>
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	// Return the router instance
	return router
}
