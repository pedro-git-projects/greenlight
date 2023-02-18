package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflit")
)

// Models will be wrap all application models
type Models struct {
	Movies MovieModel
}

// NewModels returns an initialized Models struct instance
func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}
