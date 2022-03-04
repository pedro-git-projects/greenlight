package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Creating a fixed-format JSON response from a string.
	// Raw string literals are used so as to include double-quote without needing to scape
	js := `{"status": "available", "environment": %q, "version": %q}`
	js = fmt.Sprintf(js, app.config.env, version)

	// Setting "Content-Type: application/json" in the header of our response.
	// this overrides the default "text/plain" "charset=utf-8"
	w.Header().Set("Content-Type", "application/json")

	// Write said JSON to response body
	w.Write([]byte(js))
}
