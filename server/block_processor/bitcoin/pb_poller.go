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

		// &{f58b7fc5ca63d6329acc97101e2fd66046e3ba557aa419a41061d35d2fbceddd 2 0 [{7526830bd163d916cc6f9ff7237c26bd3179935860397892f988be9f8552ab77 68   [6fef2cd291b7572220fe8cceeebbc1f4cbda9c3dd42e2d9c50e2d21ba4ded07f2b89e0f951a57497af7637e82b743d47680a009cc455dd22a44d600a04a193c0 20117f692257b2331233b5705ce9c682be8719ff1b2b64cbca290bd6faeb54423eac06a57fdf3d8901750063036f7264010118746578742f706c61696e3b636861727365743d7574662d3800357b2270223a226272632d3230222c226f70223a226d696e74222c227469636b223a2273617963222c22616d74223a2231303030227d68 c0117f692257b2331233b5705ce9c682be8719ff1b2b64cbca290bd6faeb54423e] false 4294967293}] [{512007c8867392a8c2a3a2a2c777bd71bf9b310fd7922d43e6363450559fdd13020a OP_PUSHNUM_1 OP_PUSHBYTES_32 07c8867392a8c2a3a2a2c777bd71bf9b310fd7922d43e6363450559fdd13020a v1_p2tr bc1pqlygvuuj4rp28g4zcamm6udlnvcsl4uj94p7vd352p2elhgnqg9q2xr8tl 546}] 328 610 765 {false 0  0}}

		if err != nil {
			fmt.Printf("error...")
			panic("error getting resp from poll for loop")
		}

		trans, ok := res.(bitcoin.Transaction)

		if !ok {
			fmt.Println("raw trans: ", res)
			fmt.Printf("error in pb poll for loop ...")
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
