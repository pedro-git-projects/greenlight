package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// retrieve the "id" URL parameter from teh current request context, the convert it to an integer and return.
func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

/* Defining a writeJSON() helper for sending responses.
*  It takes the destination http.ResponseWriter, the status code to send, the data to be mashaled and
*  a header map containing additional headers we might want to send in the response */
func (app *application) WriteJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	// apending newline for ease of read
	js = append(js, '\n')

	// iterating over headers map and adding to the response writer header map
	// note that iterating trough an nil map won't cause errors
	for key, value := range headers {
		w.Header()[key] = value
	}

	// Adding the Content-Type and writing headers and response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
