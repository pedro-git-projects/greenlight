package data

import (
	"time"
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
		Title:   dto.Title,
		Year:    dto.Year,
		Runtime: dto.Runtime,
		Genres:  dto.Genres,
	}
}
