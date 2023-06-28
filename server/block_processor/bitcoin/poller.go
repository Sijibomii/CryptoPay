package bitcoin

import "github.com/anthdm/hollywood/actor"

type Poller struct {
	BtcClient      *actor.PID
	PostgresClient *actor.PID
	BlockProcessor *actor.PID
	Network        string
}

func New(network string, btcClient, postgresClient, processor *actor.PID) *Poller {
	return &Poller{
		BtcClient:      btcClient,
		PostgresClient: postgresClient,
		BlockProcessor: processor,
		Network:        network,
	}
}
