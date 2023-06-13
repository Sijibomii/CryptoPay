package models

import (
	"errors"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/google/uuid"
	"github.com/sijibomii/cryptopay/types"
)

type ClientTokenPayload struct {
	ID         uuid.UUID
	Name       string
	Token      uuid.UUID
	Store_id   uuid.UUID
	Domain     string
	Created_at time.Time
	Client     types.Client
}

func NewClientToken() ClientTokenPayload {
	return ClientTokenPayload{}
}

func (sp *ClientTokenPayload) Set_created_at() error {
	sp.Created_at = time.Now()
	return nil
}

type ClientToken struct {
	ID         uuid.UUID
	Name       string
	Token      uuid.UUID
	Store_id   uuid.UUID
	Domain     string
	Created_at time.Time
	Client     types.Client
}

func (cp *ClientTokenPayload) ToClientToken() ClientToken {
	return ClientToken{
		ID:         cp.ID,
		Name:       cp.Name,
		Token:      cp.Token,
		Store_id:   cp.Store_id,
		Domain:     cp.Domain,
		Created_at: cp.Created_at,
		Client:     cp.Client,
	}
}

type InsertClientTokenMessage struct {
	Payload ClientToken
}

type FindClientTokensByStoreMessage struct {
	store_id uuid.UUID
	limit    int64
	offset   int64
}
type FindClientTokenByIdMessage struct {
	id uuid.UUID
}

type FindClientTokenByTokenAndDomainMessage struct {
	token  uuid.UUID
	domain string
}

type DeleteClientTokenMessage struct {
	id uuid.UUID
}

func InsertClientToken(e *actor.Engine, conn *actor.PID, d ClientTokenPayload) (ClientToken, error) {
	d.Set_created_at()
	var resp = e.Request(conn, InsertClientTokenMessage{
		Payload: d.ToClientToken(),
	}, 500)
	res, err := resp.Result()
	if err != nil {
		return ClientToken{}, errors.New("An error occured!")
	}
	myStruct, ok := res.(ClientToken)

	if !ok {
		return ClientToken{}, errors.New("An error occured!")
	}

	return myStruct, nil
}

func FindClientTokensByStore(e *actor.Engine, conn *actor.PID, store_id uuid.UUID, limit, offset int64) ([]ClientToken, error) {

}
