package data

import (
	"time"

	"github.com/pedro-git-projects/greenlight/internal/validator"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty,string"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}

func MovieFromDTO(dto *MovieDTO) *Movie {
	return &Movie{
		Title:   *dto.Title,
		Year:    *dto.Year,
		Runtime: *dto.Runtime,
		Genres:  dto.Genres,
	}
}

// CopyDTOFields checks for nil fields
// to permit partial updates
func (m *Movie) CopyDTOFields(dto *MovieDTO) {
	if dto.Title != nil {
		m.Title = *dto.Title
	}
	if dto.Year != nil {
		m.Year = *dto.Year
	}
	if dto.Runtime != nil {
		m.Runtime = *dto.Runtime
	}
	if dto.Genres != nil {
		m.Genres = dto.Genres
	}
}

func (movie Movie) ValidateMovie(v *validator.Validator) {
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
