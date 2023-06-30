package bitcoin

import (
	"errors"
	"fmt"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/sijibomii/cryptopay/blockchain_client/bitcoin"
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
}

func (poller *Poller) startPolling(e *actor.Engine, conn *actor.PID, ignore_prev_blocks bool) (bool, error) {

	// send bootstrap message.

	// get current block from response and send poll message

	return false, nil
}

type BootstrapPollerMessage struct {
	Ignore_previous_blocks bool
}

func (poller *Poller) bootstrapPoller()

type Poll struct {
	Block_number int
	Retry_count  int
}

func (poller *Poller) poll(e *actor.Engine, conn *actor.PID, block_number, retry_count int) (bool, error) {

	var resp = e.Request(poller.BtcClient, bitcoin.GetBlockByHeightMessage{
		Block_Height: block_number,
	}, time.Millisecond*200)

	res, err := resp.Result()
	if err != nil {
		return false, errors.New("An error occured!")
	}
	block, ok := res.(bitcoin.Block)

	if !ok {
		return false, errors.New("An error occured!")
	}

	e.Send(poller.BlockProcessor, ProcessBlockMessage{
		Block: block,
	})

	delay := time.NewTimer(time.Second * 5).C

	select {
	case <-delay:
		e.Send(conn, Poll{
			Block_number: block_number + 1,
			Retry_count:  retry_count,
		})

	default:
		return false, nil
	}

	return true, nil

}

func (poller *Poller) Receive(ctx *actor.Context) {
	switch l := ctx.Message().(type) {

	case StartPollingMessage:
		payload, err := poller.startPolling(ctx.Engine(), ctx.PID(), l.Ignore_previous_blocks)

		if err != nil {
			ctx.Respond(err.Error())
		}

		ctx.Respond(payload)

	default:
		fmt.Println("UNKNOWN MESSAGE TO POLLER CLIENT")
	}
}
