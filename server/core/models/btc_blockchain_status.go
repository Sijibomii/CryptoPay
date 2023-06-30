package models

import (
	"errors"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/google/uuid"
)

type BtcBlockChainStatusPayload struct {
	ID           uuid.UUID
	Network      string
	Block_Height int
	Created_at   time.Time
}

type BtcBlockChainStatus struct {
	ID           uuid.UUID
	Network      string
	Block_Height int
	Created_at   time.Time
}

func (bcs *BtcBlockChainStatusPayload) Set_created_at() error {
	bcs.Created_at = time.Now()
	return nil
}

func (bcs *BtcBlockChainStatusPayload) ToBtcBlockChainStatus() BtcBlockChainStatus {
	return BtcBlockChainStatus{
		ID:           bcs.ID,
		Network:      bcs.Network,
		Block_Height: bcs.Block_Height,
		Created_at:   bcs.Created_at,
	}
}

type InsertBtcBlockChainStatusMessage struct {
	Payload BtcBlockChainStatus
}

type FindBtcBlockChainStatusByNetworkMessage struct {
	Network string
}

func InsertBtcBlockChainStatus(e *actor.Engine, conn *actor.PID, d BtcBlockChainStatusPayload) (BtcBlockChainStatus, error) {
	d.Set_created_at()
	var resp = e.Request(conn, InsertBtcBlockChainStatusMessage{
		Payload: d.ToBtcBlockChainStatus(),
	}, time.Millisecond*100)
	res, err := resp.Result()
	if err != nil {
		return BtcBlockChainStatus{}, errors.New("An error occured!")
	}
	myStruct, ok := res.(BtcBlockChainStatus)

	if !ok {
		return BtcBlockChainStatus{}, errors.New("An error occured!")
	}

	return myStruct, nil
}

func FindBtcBlockChainStatusByNetwork(e *actor.Engine, conn *actor.PID, network string) (BtcBlockChainStatus, error) {
	var resp = e.Request(conn, FindBtcBlockChainStatusByNetworkMessage{
		Network: network,
	}, time.Millisecond*100)
	res, err := resp.Result()

	if err != nil {
		return BtcBlockChainStatus{}, errors.New("An error occured!")
	}
	myStruct, ok := res.(BtcBlockChainStatus)

	if !ok {
		return BtcBlockChainStatus{}, errors.New("An error occured!")
	}

	return myStruct, nil
}
