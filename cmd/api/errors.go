package main

import "net/http"

// the logError reciever is a helper for loggin error messages
func (app application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

// the errorResponse reciever is helper for sending JSON-formatted error messages
// envelope included
func (app application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// TODO: implement missing error helpers
// update functions to return propper JSON errors
