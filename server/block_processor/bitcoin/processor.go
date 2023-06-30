package bitcoin

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/sijibomii/cryptopay/blockchain_client/bitcoin"
)

type Processor struct {
	Network        string
	PostgresClient *actor.PID
}

type ProcessBlockMessage struct {
	Block bitcoin.Block
}
