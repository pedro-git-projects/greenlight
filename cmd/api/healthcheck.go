package main

import (
	"fmt"
	"net/http"
)

// Handler writing plain-text response with information about
// application status, environment setting and version
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	// not accessing trough the application struct because its currently hardcoded
	fmt.Fprintf(w, "version: %s\n", version)

}
