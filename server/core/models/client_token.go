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
	var resp = e.Request(conn, FindClientTokensByStoreMessage{
		store_id: store_id,
		limit:    limit,
		offset:   offset,
	}, 500)
	res, err := resp.Result()
	var cts []ClientToken
	if err != nil {
		return cts, errors.New("An error occured!")
	}
	myStruct, ok := res.([]ClientToken)

	if !ok {
		return cts, errors.New("An error occured!")
	}

	return myStruct, nil
}

func FindClientTokenById(e *actor.Engine, conn *actor.PID, id uuid.UUID) (ClientToken, error) {
	var resp = e.Request(conn, FindClientTokenByIdMessage{
		id: id,
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

func FindClientTokenByTokenAndDomain(e *actor.Engine, conn *actor.PID, token uuid.UUID, domain string) (ClientToken, error) {
	var resp = e.Request(conn, FindClientTokenByTokenAndDomainMessage{
		token:  token,
		domain: domain,
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

func DeleteClientToken(e *actor.Engine, conn *actor.PID, id uuid.UUID) (bool, error) {
	var resp = e.Request(conn, DeleteClientTokenMessage{
		id: id,
	}, 500)
	res, err := resp.Result()
	if err != nil {
		return false, errors.New("An error occured!")
	}
	myStruct, ok := res.(bool)

	if !ok {
		return false, errors.New("An error occured!")
	}

	return myStruct, nil
}
