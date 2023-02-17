package data

import "database/sql"

// MovieModel defines a type that wraps a sql.DB connection pool
type MovieModel struct {
	DB *sql.DB
}

// Insert inserts movies on the movies table
func (m MovieModel) Insert(movie *Movie) error {
	return nil
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
