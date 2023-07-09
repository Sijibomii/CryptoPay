package util

import (
	"crypto/rsa"
	"net/http"

	"github.com/anthdm/hollywood/actor"
	"github.com/sijibomii/cryptopay/config"
	"github.com/sijibomii/cryptopay/core/utils"
)

type AppState struct {
	PgExecutor      utils.PgExecutor
	Config          *config.Config
	Postgres        *actor.PID
	Engine          *actor.Engine
	Mailer          *actor.PID
	CoinClient      *actor.PID
	BtcClient       *actor.PID
	ProcessorClient *actor.PID
	PollerClient    *actor.PID
	PBPollerClient  *actor.PID
	PrivateKey      *rsa.PrivateKey
}

type AppHandler struct {
	*AppState
	HandlerFunc http.HandlerFunc
}
