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

	var resp = e.Request(pb.BtcClient, bitcoin.GetRawMempoolMessage{}, time.Millisecond*200)

	res, err := resp.Result()
	if err != nil {
		fmt.Printf("error...")
		panic("error getting resp from boostrap")
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
		}, time.Millisecond*200)

		res, err := resp.Result()

		if err != nil {
			fmt.Printf("error...")
			panic("error getting resp from poll for loop")
		}

		trans, ok := res.(bitcoin.Transaction)

		if !ok {
			fmt.Printf("error in pb poll for loop ...")

		}

		transResult = append(transResult, trans)
	}

	e.Send(pb.BlockProcessor, ProcessMempoolTransactionsMessage{
		Transactions: transResult,
	})

	time.Sleep(4 * time.Second)

	e.Send(conn, PBPollMessage{
		Previous:    mempool,
		Retry_count: retry_size,
	})

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