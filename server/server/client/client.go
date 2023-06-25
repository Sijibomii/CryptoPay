package client

import (
	"time"

	"github.com/google/uuid"
	"github.com/sijibomii/cryptopay/core/models"
)

type Client struct {
	ID         uuid.UUID
	Store_id   uuid.UUID
	Created_at time.Time
}

func NewClient(clientToken models.ClientToken) Client {
	return Client{
		ID:         uuid.New(),
		Store_id:   clientToken.Store_id,
		Created_at: time.Now(),
	}
}
