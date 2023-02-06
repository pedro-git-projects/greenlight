package main

import "github.com/pedro-git-projects/greenlight/internal/data"

type movieDTO struct {
	Title   string       `json:"title"`
	Year    int32        `json:"year"`
	Runtime data.Runtime `json:"runtime"`
	Genres  []string     `json:"genres"`
}
