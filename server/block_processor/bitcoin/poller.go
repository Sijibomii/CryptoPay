package bitcoin

import (
	"errors"
	"fmt"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/google/uuid"
	"github.com/sijibomii/cryptopay/blockchain_client/bitcoin"
	"github.com/sijibomii/cryptopay/core/models"
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
	fmt.Print("START POLLING MESSAGE \n")
	// send bootstrap message.

	current_block, err := poller.bootstrapPoller(e, conn, ignore_prev_blocks)
	// var resp = e.Request(conn, BootstrapPollerMessage{
	// 	Ignore_previous_blocks: true,
	// }, time.Millisecond*200)

	// _, err := resp.Result()
	if err != nil {
		fmt.Printf("error...123 %s", err.Error())
		panic("error getting resp from boostrap")
	}

	// currentblock, ok := res.(int)

	fmt.Print("CURRENT BLOCK  ", current_block, "\n")

	// if !ok {
	// 	fmt.Printf("error...")
	// 	panic("error parsing resp from bootstrap")
	// }

	// get current block from response and send poll message

	fmt.Print("POLL MESSAGE SENDING... \n")
	e.Send(conn, Poll{
		Block_number: current_block,
		Retry_count:  0,
	})

	return false, nil
}

type BootstrapPollerMessage struct {
	Ignore_previous_blocks bool
}

func (poller *Poller) bootstrapPoller(e *actor.Engine, conn *actor.PID, ig_prev_blocks bool) (int, error) {
	fmt.Printf("BLOCK COUNT MESSAGE SENT \n")
	var resp = e.Request(poller.BtcClient, bitcoin.GetBlockCountMessage{}, time.Millisecond*1000)

	res, err := resp.Result()
	if err != nil {
		fmt.Printf("error... \n")
		panic("error getting block count")
	}

	blockCount, ok := res.(int)

	if !ok {
		fmt.Printf("error...")
		panic("error getting block count")
	}

	status, err := models.FindBtcBlockChainStatusByNetwork(e, poller.PostgresClient, poller.Network)

	if status.ID == uuid.Nil || status.Network == "" || status.Block_Height == 0 || status.Created_at.IsZero() {
		// insert new one
		payload := models.BtcBlockChainStatusPayload{
			Network:      poller.Network,
			Block_Height: blockCount,
		}

		status, err = models.InsertBtcBlockChainStatus(e, poller.PostgresClient, payload)
	}

	if err != nil {
		fmt.Printf("error...")
		panic("error finding or inserting bloack status")
	}

	return status.Block_Height, nil

}

type Poll struct {
	Block_number int
	Retry_count  int
}

func (poller *Poller) poll(e *actor.Engine, conn *actor.PID, block_number, retry_count int) (bool, error) {

	var resp = e.Request(poller.BtcClient, bitcoin.GetBlockByHeightMessage{
		Block_Height: block_number,
	}, time.Millisecond*1200)

	res, err := resp.Result()

	if err != nil {
		return false, errors.New("An error occured!")
	}
	str, ok := res.(string)

	if !ok {
		return false, errors.New("An error occured!")
	}
	block, err := parseBlockString(str)

	e.Send(poller.BlockProcessor, ProcessBlockMessage{
		Block: block,
	})
	fmt.Print("hereee sending poll2")
	delay := time.NewTimer(time.Second * 5).C

	select {
	case <-delay:
		fmt.Print("hereee sending poll")
		e.Send(conn, Poll{
			Block_number: block_number + 1,
			Retry_count:  retry_count,
		})

	default:
		return false, nil
	}

	return true, nil

}

func parseBlockString(str string) (bitcoin.Block, error) {
	var block bitcoin.Block
	_, err := fmt.Sscanf(str, "Block: ID=%s, Height=%d, Version=%d, Timestamp=%d, TxCount=%d, Size=%d, Weight=%d, MerkleRoot=%s, PreviousBlock=%s, MedianTime=%d, Nonce=%d, Bits=%d, Difficulty=%d",
		&block.ID, &block.Height, &block.Version, &block.Timestamp, &block.TxCount, &block.Size, &block.Weight, &block.MerkleRoot, &block.PreviousBlockHash, &block.MedianTime, &block.Nonce, &block.Bits, &block.Difficulty)
	if err != nil {
		return bitcoin.Block{}, err
	}
	return block, nil
}

func (poller *Poller) Receive(ctx *actor.Context) {
	switch l := ctx.Message().(type) {

	case StartPollingMessage:

		_, _ = poller.startPolling(ctx.Engine(), ctx.PID(), l.Ignore_previous_blocks)

		// if err != nil {
		// 	ctx.Respond(err.Error())
		// }

		// ctx.Respond(payload)

	case Poll:
		poller.poll(ctx.Engine(), ctx.PID(), l.Block_number, l.Retry_count)

	default:
		fmt.Println("UNKNOWN MESSAGE TO POLLER CLIENT")
	}
}
