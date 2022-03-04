package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pedro-git-projects/greenlight/internal/data"
)

// placeholder for movie creation
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
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
	err = app.WriteJSON(w, http.StatusOK, movie, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

}
