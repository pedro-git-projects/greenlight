package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/pedro-git-projects/greenlight/internal/data"
	"github.com/pedro-git-projects/greenlight/internal/validator"
)

// createMovieHandler reads data from the client and creates a new movie
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	input := &data.MovieDTO{}
	if err := app.readJSON(w, r, &input); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	if data.ValidateMovieDTO(v, input); !v.IsValid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	m := data.MovieFromDTO(input)
	if err := app.models.Movies.Insert(m); err != nil {
		app.internalServerErr(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("v1/movies/%d", m.ID))

	if err := app.writeJSON(w, http.StatusCreated, envelope{"movie": m}, headers); err != nil {
		app.internalServerErr(w, r, err)
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

	movie, err := app.models.Movies.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.internalServerErr(w, r, err)
		}
		return
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil); err != nil {
		app.internalServerErr(w, r, err)
	}
}

func (app *application) updateMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie, err := app.models.Movies.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.internalServerErr(w, r, err)
		}
		return
	}

	input := &data.MovieDTO{}
	if err := app.readJSON(w, r, &input); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	movie.CopyDTOFields(input)
	v := validator.New()
	if movie.ValidateMovie(v); !v.IsValid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	if err := app.models.Movies.Update(movie); err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.internalServerErr(w, r, err)
		}
		return
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil); err != nil {
		app.internalServerErr(w, r, err)
	}
}

func (app *application) deleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	if err := app.models.Movies.Delete(id); err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.internalServerErr(w, r, err)
		}
	}

	err = app.writeJSON(w, http.StatusNoContent, envelope{"message": "movie successfully deleted"}, nil)
	if err != nil {
		app.internalServerErr(w, r, err)
	}

}

func (app *application) listMoviesHandler(w http.ResponseWriter, r *http.Request) {

	input := data.EmptyMovieDTO()
	input.Filters.SortSafeList = []string{"id", "title", "year", "runtime", "-id", "-title", "-year", "-runtime"}

	v := validator.New()
	qs := r.URL.Query()

	input.Genres = app.readCSV(qs, "genres", []string{})

	title := app.readString(qs, "title", "")
	input.Title = &title

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Filters.Sort = app.readString(qs, "sort", "id")

	if data.ValidateFilters(v, input.Filters); !v.IsValid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	movies, metadata, err := app.models.Movies.GetAll(*input.Title, input.Genres, input.Filters)
	if err != nil {
		app.internalServerErr(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movies": movies, "metadata": metadata}, nil)
	if err != nil {
		app.internalServerErr(w, r, err)
	}
}
