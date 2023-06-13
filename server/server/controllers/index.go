package controllers

import (
	"fmt"
	"net/http"

	"github.com/sijibomii/cryptopay/server/util"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, appState *util.AppState) {
	// Access the appState in your handler function
	// You can use appState.Postgres, appState.Mailer, appState.Config, etc.
	fmt.Fprintf(w, "Hello, world! Database: %v", appState.Postgres)
}
