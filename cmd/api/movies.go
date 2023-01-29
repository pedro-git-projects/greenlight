package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pedro-git-projects/greenlight/internal/data"
)

// createMovieHandler reads data from the client and creates a new movie
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	input := &movieDTO{}
	if err := json.NewDecoder(r.Body).Decode(input); err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Fprintf(w, "%v\n", *input)
}

// showMovieHandler displays the movie with specified ID
// it redirects to a 404 otherwise
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := &data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"romance", "drama", "war"},
		Version:   1,
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil); err != nil {
		app.internalServerErr(w, r, err)
	}

}
