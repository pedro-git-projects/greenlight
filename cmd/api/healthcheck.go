package main

import (
	"net/http"
)

// healthcheckHandler is used to send the availability, environment and version of the application as JSON to the client
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"healthcheck": data}, nil); err != nil {
		app.internalServerErr(w, r, err)
	}
}
