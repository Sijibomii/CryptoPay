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
		ctx.Respond(payload)

	case GetRawTransactionMessage:
		payload, err := client.get_Transaction_By_Hash(l.Transaction_Hash)

		if err != nil {
			ctx.Respond(err.Error())
		}
		ctx.Respond(payload)

	case GetAllTransactionsByBlockHeightMessage:

		payload, err := client.get_all_transactions_by_block_height(l.Block_Height)

		if err != nil {
			ctx.Respond(err.Error())
		}
		ctx.Respond(payload)

	case BroadcastRawTransactionMessage:

		payload := client.broadcastTransaction(l.RawTx)

		ctx.Respond(payload)

	case GetFeeEstimatesMessage:
		payload, err := client.getFeeEstimates()

		if err != nil {
			ctx.Respond(err.Error())
		}
		ctx.Respond(payload)

	default:
		fmt.Println("UNKNOWN MESSAGE TO BLOACKCHAIN CLIENT")
	}
}

func GetBlockCount(e *actor.Engine, conn *actor.PID) (int, error) {

	var resp = e.Request(conn, GetBlockCountMessage{}, time.Millisecond*100)

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
	}, time.Millisecond*100)

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
	}, time.Millisecond*100)

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
	}, time.Millisecond*100)

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

func BroadcastRawTransaction(e *actor.Engine, conn *actor.PID, rawTx string) (bool, error) {
	var resp = e.Request(conn, BroadcastRawTransactionMessage{
		RawTx: rawTx,
	}, time.Millisecond*100)

	res, err := resp.Result()
	if err != nil {
		return false, errors.New("An error occured!")
	}

	error, ok := res.(error)

	if !ok {
		return true, nil
	}

	return false, error
}

func GetFeeEstimates(e *actor.Engine, conn *actor.PID) (*FeeEstimates, error) {
	var resp = e.Request(conn, GetFeeEstimatesMessage{}, time.Millisecond*100)

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
	var resp = e.Request(conn, GetRawMempoolMessage{}, time.Millisecond*100)

	res, err := resp.Result()
	if err != nil {
		return nil, errors.New("An error occured!")
	}

	mempool, ok := res.([]MempoolEntry)

	if !ok {
		return nil, errors.New("An error occured!")
	}

	return &mempool, nil
}

func GetAllTransactionsByBlockHeight(e *actor.Engine, conn *actor.PID, block_height int) ([]Transaction, error) {

	var resp = e.Request(conn, GetAllTransactionsByBlockHeightMessage{
		Block_Height: block_height,
	}, time.Millisecond*500)

	res, err := resp.Result()

	if err != nil {
		return nil, errors.New("An error occured!")
	}

	trans, ok := res.([]Transaction)

	if !ok {
		return nil, errors.New("An error occured!")
	}

	return trans, nil
}
