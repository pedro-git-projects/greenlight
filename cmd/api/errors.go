package main

import (
	"fmt"
	"net/http"
)

// the logError reciever is a helper for loggin error messages
func (app *application) logError(r *http.Request, err error) {
	app.logger.PrintError(err, nil)
}

// the errorResponse reciever is helper for sending JSON-formatted error messages
// envelope included
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// internalServerErr is used to send a 500 status code and JSON response to the client
func (app *application) internalServerErr(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	msg := "The server encountered an error and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, msg)
}

// notFoundResponse is used to send a 404 status code and JSON response to the client
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	msg := "Resource not found"
	app.errorResponse(w, r, http.StatusNotFound, msg)
}

// methodNotAllowedResponse  used to send a 405 status code and JSON response to the client
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("Method %s is not allowed for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, msg)
}

// badRequestResponse is halper moethod for sending a bad request client response
func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func (app *application) editConflictResponse(w http.ResponseWriter, r *http.Request) {
	message := "unable to update the record due to an edit conflict, please try again"
	app.errorResponse(w, r, http.StatusConflict, message)
}
