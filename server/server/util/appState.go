package util

import (
	"net/http"

	"github.com/anthdm/hollywood/actor"
	coinclient "github.com/sijibomii/cryptopay/coin_client"
	"github.com/sijibomii/cryptopay/config"
	"github.com/sijibomii/cryptopay/core/utils"
)

type AppState struct {
	PgExecutor utils.PgExecutor
	Config     *config.Config
	Postgres   *actor.PID
	Engine     *actor.Engine
	Mailer     *actor.PID
	CoinClient *coinclient.CoinClient
}

type AppHandler struct {
	*AppState
	HandlerFunc http.HandlerFunc
}
