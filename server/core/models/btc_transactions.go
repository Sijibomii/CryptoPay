package models

import (
	"errors"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/google/uuid"
	"github.com/sijibomii/cryptopay/blockchain_client/bitcoin"
)

type BtcTransaction struct {
	ID          uuid.UUID           `json:"id"`
	Hash        string              `json:"hash"`
	Transaction bitcoin.Transaction `json:"transaction"`
}

type BtcTransactionPayload struct {
	ID          uuid.UUID
	Hash        string
	Transaction bitcoin.Transaction
}

func (p *BtcTransactionPayload) Set_id() error {
	p.ID = uuid.New()
	return nil
}

type InsertBtcTransactionMessage struct {
	Payload BtcTransaction
}

func (p *BtcTransactionPayload) ToBtcTransaction() BtcTransaction {
	return BtcTransaction{
		ID:          p.ID,
		Hash:        p.Hash,
		Transaction: p.Transaction,
	}
}

func InsertTransaction(e *actor.Engine, conn *actor.PID, d BtcTransactionPayload) (BtcTransaction, error) {
	d.Set_id()
	var resp = e.Request(conn, InsertBtcTransactionMessage{
		Payload: d.ToBtcTransaction(),
	}, time.Millisecond*100)

	res, err := resp.Result()
	if err != nil {
		return BtcTransaction{}, errors.New("An error occured!")
	}
	myStruct, ok := res.(BtcTransaction)

	if !ok {
		return BtcTransaction{}, errors.New("An error occured!")
	}

	return myStruct, nil
}
