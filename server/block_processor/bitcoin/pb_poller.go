package bitcoin

import (
	"fmt"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/sijibomii/cryptopay/blockchain_client/bitcoin"
)

type PBPoller struct {
	BtcClient      *actor.PID
	PostgresClient *actor.PID
	BlockProcessor *actor.PID
	Network        string
}

func NewPBP(network string, btcClient, postgresClient, processor *actor.PID) *PBPoller {
	return &PBPoller{
		BtcClient:      btcClient,
		PostgresClient: postgresClient,
		BlockProcessor: processor,
		Network:        network,
	}
}

type StartPBPollingMessage struct{}

func (pb *PBPoller) startPolling(e *actor.Engine, conn *actor.PID) {

	var mem []bitcoin.MempoolEntry

	e.Send(conn, PBPollMessage{
		Previous:    mem,
		Retry_count: 0,
	})

}

type PBPollMessage struct {
	Previous    []bitcoin.MempoolEntry
	Retry_count int
}

func (pb *PBPoller) poll(e *actor.Engine, conn *actor.PID, previous []bitcoin.MempoolEntry, retry_size int) {
	fmt.Print("\n POLLING FROM PRENDING BLOCK POLLER ######### \n")

	var resp = e.Request(pb.BtcClient, bitcoin.GetRawMempoolMessage{}, time.Millisecond*2000)

	res, err := resp.Result()
	if err != nil {
		fmt.Printf("error...")
		panic(err.Error())
	}

	mempool, ok := res.([]bitcoin.MempoolEntry)

	if !ok {
		fmt.Printf("error...")
		panic("error parsing resp from poll in pb")
	}

	diff := difference(mempool, previous)

	var transResult []bitcoin.Transaction

	for _, transaction := range diff {

		var resp = e.Request(pb.BtcClient, bitcoin.GetRawTransactionMessage{
			Transaction_Hash: transaction.TxID,
		}, time.Millisecond*2000)

		res, err := resp.Result()

		if err != nil {
			fmt.Printf("error...")
			panic("error getting resp from poll for loop")
		}

		trans, ok := res.(bitcoin.Transaction)

		if !ok || trans.TxID == "0" {
			fmt.Println("raw trans: ", res)
			fmt.Printf("error in pb poll for loop ...")
			continue
		}

		transResult = append(transResult, trans)

		fmt.Print("\n RAW TRANSACTION APPENDED FROM PB POLLER: ", trans)
		fmt.Println("")
	}

	e.Send(pb.BlockProcessor, ProcessMempoolTransactionsMessage{
		Transactions: transResult,
	})

	e.SendRepeat(conn, PBPollMessage{
		Previous:    mempool,
		Retry_count: retry_size,
	}, time.Millisecond*4000)

}

func difference(sliceA, sliceB []bitcoin.MempoolEntry) []bitcoin.MempoolEntry {
	result := make([]bitcoin.MempoolEntry, 0)

	// Create a map to efficiently check if an element exists in slice B
	setB := make(map[string]struct{})
	for _, element := range sliceB {
		setB[element.TxID] = struct{}{}
	}

	// Iterate over slice A and check if each element is in slice B
	for _, element := range sliceA {
		if _, ok := setB[element.TxID]; !ok {
			result = append(result, element)
		}
	}

	return result
}

// messages receive
func (poller *PBPoller) Receive(ctx *actor.Context) {
	switch l := ctx.Message().(type) {

	case actor.Started:
		fmt.Println("pending poller actor started")

	case StartPBPollingMessage:
		fmt.Println("START POLLING FROM PENDING BLOCK POLLER")
		poller.startPolling(ctx.Engine(), ctx.PID())

	case PBPollMessage:
		poller.poll(ctx.Engine(), ctx.PID(), l.Previous, l.Retry_count)

	default:
		fmt.Println("UNKNOWN MESSAGE TO PBPoller CLIENT")
	}
}
