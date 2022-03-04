package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`                // unique movie ID
	CreatedAt time.Time `json:"-"`                 // "-" will hide this field // timestamp for movie creation
	Title     string    `json:"title,omitempty"`   // title
	Year      int32     `json:"year,omitempty"`    // release year
	Runtime   int32     `json:"runtime,omitempty"` // runtime in minutes
	Genres    []string  `json:"genres,omitempty"`  // slice of genres
	Version   int32     `json:"version"`           // starts with 1 and will be incremented each time the movie is updated
}
