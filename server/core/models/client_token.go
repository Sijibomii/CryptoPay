package models

import (
	"time"

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
