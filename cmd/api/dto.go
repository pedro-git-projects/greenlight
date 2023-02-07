package main

import (
	"time"

	"github.com/pedro-git-projects/greenlight/internal/data"
	"github.com/pedro-git-projects/greenlight/internal/validator"
)

const (
	required = "must be provided"
)

type movieDTO struct {
	Title   string       `json:"title"`
	Year    int32        `json:"year"`
	Runtime data.Runtime `json:"runtime"`
	Genres  []string     `json:"genres"`
}

func validateMovieDTO(v *validator.Validator, movie *movieDTO) {
	v.Check(movie.Title != "", "title", required)
	v.Check(len(movie.Title) <= 500, "title", "must be nomore than 500 bytes long")

	v.Check(movie.Year != 0, "year", required)
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(movie.Runtime != 0, "runtime", required)
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(movie.Genres != nil, "genres", required)
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must contain at most 5 genres")
	v.Check(validator.UniqueKeys(movie.Genres), "genres", "must not contain duplicate values")
}
