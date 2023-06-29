package bitcoin

import (
	"fmt"

	"github.com/anthdm/hollywood/actor"
)

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

type StartPollingMessage struct {
	Ignore_previous_blocks bool
	Engine                 *actor.Engine
	Pid                    *actor.PID
}

func startPolling(self *Poller, e *actor.Engine, conn *actor.PID, ignore_prev_blocks bool) (bool, error) {
	return false, nil
}

type BootstrapPollerMessage struct {
	Ignore_previous_blocks bool
}

func (poller *Poller) Receive(ctx *actor.Context) {
	switch l := ctx.Message().(type) {

	default:
		fmt.Println("UNKNOWN MESSAGE TO POLLER CLIENT")
	}
}
