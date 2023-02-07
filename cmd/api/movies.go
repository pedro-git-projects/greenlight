package main

import (
	"net/http"
	"time"

	"github.com/pedro-git-projects/greenlight/internal/data"
	"github.com/pedro-git-projects/greenlight/internal/validator"
)

// createMovieHandler reads data from the client and creates a new movie
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	input := &movieDTO{}
	if err := app.readJSON(w, r, &input); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	if validateMovieDTO(v, input); !v.IsValid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

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
