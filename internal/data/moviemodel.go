package data

import (
	"database/sql"

	"github.com/lib/pq"
)

// MovieModel defines a type that wraps a sql.DB connection pool
type MovieModel struct {
	DB *sql.DB
}

// Insert inserts movies on the movies table
func (m MovieModel) Insert(movie *Movie) error {
	query := `
	INSERT INTO movies (title, year, runtime, genres)
	VALUES($1, $2, $3, $4)
	RETURNING id, created_at, version`

	args := []any{movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres)}

	return m.DB.QueryRow(query, args...).Scan(&movie.ID, &movie.CreatedAt, &movie.Version)
}

// Get fetches a specific record from the movies table
func (m MovieModel) Get(movie *Movie) error {
	return nil
}

// Update updates a specific record from the movies table
func (m MovieModel) Update(movie *Movie) error {
	return nil
}

// Delete deletes a specific record from the movies table
func (m MovieModel) Delete(id int64) error {
	return nil
}
