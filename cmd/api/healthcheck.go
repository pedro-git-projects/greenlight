package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// map that holds all information that will be sent in the response
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	// Marshaling the map to JSON.
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}

	// appending a newline is just cosmetic, it will make it easier to read in the terminal
	js = append(js, '\n')

	// set content type and write response
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
