package bitcoin

import (
	"errors"
	"fmt"
	"time"

	"github.com/anthdm/hollywood/actor"
)

type GetBlockCountMessage struct{}

type GetBlockMessage struct {
	Block_hash string
}

type GetBlockByHeightMessage struct {
	Block_Height int
}

type GetRawTransactionMessage struct {
	Transaction_Hash string
}

type BroadcastRawTransactionMessage struct {
	RawTx string
}

type GetFeeEstimatesMessage struct{}

type GetRawMempoolMessage struct{}

type GetAllTransactionsByBlockHeightMessage struct {
	Block_Height int
}

func (client *BlockchainClient) Receive(ctx *actor.Context) {
	switch l := ctx.Message().(type) {

	case GetBlockCountMessage:
		//fmt.Printf("revcieved block count message \n")

		payload, err := client.get_block_count()

		if err != nil {
			ctx.Respond(err.Error())
		}
		ctx.Respond(payload)

	case GetBlockMessage:
		payload, err := client.get_Block(l.Block_hash)

		if err != nil {
			ctx.Respond(err.Error())
		}
		ctx.Respond(payload)

	case GetBlockByHeightMessage:

		payload, err := client.get_Block_with_height(l.Block_Height)

		if err != nil {
			ctx.Respond(err.Error())
		}
		ctx.Respond(payload.String())

	case GetRawTransactionMessage:
		payload, err := client.get_Transaction_By_Hash(l.Transaction_Hash)

		if err != nil {
			ctx.Respond(err.Error())
		}

		// stringify:
		//fmt.Print("\n TRANSACTION RAW MESSAGE : ", *payload)
		//fmt.Println("")
		ctx.Respond(*payload)

	case GetAllTransactionsByBlockHeightMessage:

		payload, err := client.get_all_transactions_by_block_height(l.Block_Height)

		fmt.Print("\n payload equals: ", payload)

		if err != nil {
			fmt.Printf("error: %s", err.Error())
			fmt.Println("ERRORRR!!!")
			ctx.Respond(err.Error())
		}

		ctx.Respond(payload)

	case BroadcastRawTransactionMessage:

		payload, err := client.broadcastTransaction(l.RawTx)

		if err != nil {
			ctx.Respond(err.Error())
		}

		ctx.Respond(payload)

	case GetFeeEstimatesMessage:
		payload, err := client.getFeeEstimates()

		if err != nil {
			ctx.Respond(err.Error())
		}
		ctx.Respond(payload)

	case GetRawMempoolMessage:
		payload, err := client.GetRawMempool()

		if err != nil {
			ctx.Respond(err.Error())
		}
		ctx.Respond(payload)

	default:
		//fmt.Print("\n UNKNOWN MESSAGE TO BLOACKCHAIN CLIENT: ", ctx.Message())
	}
}

func GetBlockCount(e *actor.Engine, conn *actor.PID) (int, error) {

	var resp = e.Request(conn, GetBlockCountMessage{}, time.Millisecond*1000)

	res, err := resp.Result()
	if err != nil {
		return 0, errors.New("An error occured!")
	}
	block_number, ok := res.(int)

	if !ok {
		return 0, errors.New("An error occured!")
	}

	return block_number, nil
}

func GetBlock(e *actor.Engine, conn *actor.PID, block_hash string) (*Block, error) {
	var resp = e.Request(conn, GetBlockMessage{
		Block_hash: block_hash,
	}, time.Millisecond*1000)

	res, err := resp.Result()
	if err != nil {
		return nil, errors.New("An error occured!")
	}
	block, ok := res.(Block)

	if !ok {
		return nil, errors.New("An error occured!")
	}

	return &block, nil
}

func GetBlockByHeight(e *actor.Engine, conn *actor.PID, block_height int) (*Block, error) {
	var resp = e.Request(conn, GetBlockByHeightMessage{
		Block_Height: block_height,
	}, time.Millisecond*1000)

	res, err := resp.Result()
	if err != nil {
		return nil, errors.New("An error occured!")
	}
	block, ok := res.(Block)

	if !ok {
		return nil, errors.New("An error occured!")
	}

	return &block, nil
}

func GetRawTransaction(e *actor.Engine, conn *actor.PID, transaction_hash string) (*Transaction, error) {

	var resp = e.Request(conn, GetRawTransactionMessage{
		Transaction_Hash: transaction_hash,
	}, time.Millisecond*1000)

	res, err := resp.Result()
	if err != nil {
		return nil, errors.New("An error occured!")
	}
	trans, ok := res.(Transaction)

	if !ok {
		return nil, errors.New("An error occured!")
	}

	return &trans, nil

}

func BroadcastRawTransaction(e *actor.Engine, conn *actor.PID, rawTx string) (string, error) {
	var resp = e.Request(conn, BroadcastRawTransactionMessage{
		RawTx: rawTx,
	}, time.Millisecond*1000)

	res, err := resp.Result()

	if err != nil {
		return "", errors.New("An error occured!")
	}

	str, ok := res.(string)

	if !ok {
		return "", errors.New("An error occured!")
	}

	return str, nil
}

func GetFeeEstimates(e *actor.Engine, conn *actor.PID) (*FeeEstimates, error) {
	var resp = e.Request(conn, GetFeeEstimatesMessage{}, time.Millisecond*1000)

	res, err := resp.Result()
	if err != nil {
		return nil, errors.New("An error occured!")
	}

	fee, ok := res.(FeeEstimates)

	if !ok {
		return nil, errors.New("An error occured!")
	}

	return &fee, nil
}

func GetRawMempool(e *actor.Engine, conn *actor.PID) (*[]MempoolEntry, error) {
	var resp = e.Request(conn, GetRawMempoolMessage{}, time.Millisecond*2000)

	res, err := resp.Result()

	//fmt.Printf("\n res returned mempool \n")

	if err != nil {
		return nil, errors.New("An error occured trying to get res!")
	}

	//fmt.Printf("\n res returned mempool .....\n")

	mempool, ok := res.([]MempoolEntry)

	if !ok {
		return nil, errors.New("An error occured trying to conv to mempool!")
	}

	return &mempool, nil
}

func GetAllTransactionsByBlockHeight(e *actor.Engine, conn *actor.PID, block_height int) ([]Transaction, error) {

	var resp = e.Request(conn, GetAllTransactionsByBlockHeightMessage{
		Block_Height: block_height,
	}, time.Millisecond*1200)

	res, err := resp.Result()

	if err != nil {
		return nil, errors.New(err.Error())
	}

	trans, ok := res.([]Transaction)

	if !ok {
		return nil, errors.New("An error occured tying to conv to transaction")
	}

	return trans, nil
}
