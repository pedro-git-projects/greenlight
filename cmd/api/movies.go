package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// placeholder for movie creation
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	// When httprouter is parsing a request, any interpolated URL parameters will be
	// stored in the request context. We can use the ParamsFromContext() function to
	// retrive a slice containing these parameters names and values
	params := httprouter.ParamsFromContext(r.Context())

	// using the ByName() method to get the value of the "id" parameter from the slice
	// we test the validity of the parameters, if they fail the test, a 404 response is written
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// Otherwise, interpolate the movie ID in a placeholder response.
	fmt.Fprintf(w, "show the details of movie %d\n", id)
}
