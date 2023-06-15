package util

import (
	"net/http"

	"github.com/anthdm/hollywood/actor"
	"github.com/sijibomii/cryptopay/config"
	"github.com/sijibomii/cryptopay/core/utils"
)

type AppState struct {
	PgExecutor utils.PgExecutor
	Config     *config.Config
	Postgres   *actor.PID
	Engine     *actor.Engine
}

type AppHandler struct {
	*AppState
	HandlerFunc http.HandlerFunc
}
