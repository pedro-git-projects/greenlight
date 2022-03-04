package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pedro-git-projects/greenlight/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	// Anonymous struct to hold the information we expect to recieve in the request body
	// Its fields are a subset of our Movie struct
	// This struct will be our target decode destination
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	// Decoding from http to struct
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	// Dumping the contents of the input struct into an HTTP response
	fmt.Fprintf(w, "%v+\n", input)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// Dummy movie utlizing the id extracted from the URL
	movie := data.Movie{
		ID:        id,
		Title:     "Taxi Driver",
		CreatedAt: time.Now(),
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	// Marshaling to json and sending as response
	err = app.WriteJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
