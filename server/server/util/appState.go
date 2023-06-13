package util

import (
	"net/http"

	"github.com/sijibomii/cryptopay/config"
	"github.com/sijibomii/cryptopay/core/utils"
)

type AppState struct {
	Postgres utils.PgExecutor
	Config   *config.Config
}

type AppHandler struct {
	*AppState
	HandlerFunc http.HandlerFunc
}
