package data

import (
	"time"

	"github.com/pedro-git-projects/greenlight/internal/validator"
)

const (
	required = "must be provided"
)

// All dto fields are pointer because
// the zero value of a pointer type is nil.
// This facilitates user input checking
type MovieDTO struct {
	Title   *string  `json:"title"`
	Year    *int32   `json:"year"`
	Runtime *Runtime `json:"runtime"`
	Genres  []string `json:"genres"`
	Filters
}

func NewMovieDTO(title *string, year *int32, runtime *Runtime, genres []string, page, pageSize int, sort string) *MovieDTO {
	return &MovieDTO{
		Title:   title,
		Year:    year,
		Runtime: runtime,
		Genres:  genres,
		Filters: *NewEmptyFilters(),
	}
}

func EmptyMovieDTO() *MovieDTO {
	r := new(int32)
	runtime := Runtime(*r)
	return &MovieDTO{
		Title:   new(string),
		Year:    new(int32),
		Runtime: &runtime,
		Genres:  make([]string, 0),
		Filters: *NewEmptyFilters(),
	}
}

func ValidateMovieDTO(v *validator.Validator, movie *MovieDTO) {
	v.Check(*movie.Title != "", "title", required)
	v.Check(len(*movie.Title) <= 500, "title", "must be nomore than 500 bytes long")

	v.Check(*movie.Year != 0, "year", required)
	v.Check(*movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(*movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(*movie.Runtime != 0, "runtime", required)
	v.Check(*movie.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(movie.Genres != nil, "genres", required)
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must contain at most 5 genres")
	v.Check(validator.UniqueKeys(movie.Genres), "genres", "must not contain duplicate values")
}
