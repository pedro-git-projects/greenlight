package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// the envelope type is used to create a top level wrapper around JSON responses
type envelope map[string]any

// readIDParam reads the route parameter id and erros out if it nil or less than 1
func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id")
	}

	return id, nil
}

// writeJSON takes a writer, status, enveloped data and header.
// It mashalls the data to JSON, if possible, else it erros out
// That done it sets the passed headers and sets the contant type to JSON
// Finally it writes the headers and sends the JSON response
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	json, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	json = append(json, '\n') // adds new line for readability

	// note that http.Header is of type map[string][]string
	// so here we iterate through this map
	// adding each header to the actual response header
	for k, v := range headers {
		w.Header()[k] = v
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json)

	return nil
}
