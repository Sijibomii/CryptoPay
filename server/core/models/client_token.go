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
	Store_id uuid.UUID
	Limit    int64
	Offset   int64
}
type FindClientTokenByIdMessage struct {
	Id uuid.UUID
}

type FindClientTokenByTokenAndDomainMessage struct {
	Token  uuid.UUID
	Domain string
}

type DeleteClientTokenMessage struct {
	Id uuid.UUID
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
		Store_id: store_id,
		Limit:    limit,
		Offset:   offset,
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
		Id: id,
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
		Token:  token,
		Domain: domain,
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
		Id: id,
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
