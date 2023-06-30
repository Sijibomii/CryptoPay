package models

import (
	"github.com/google/uuid"
	"github.com/sijibomii/cryptopay/blockchain_client/bitcoin"
)

type BtcTransaction struct {
	ID          uuid.UUID           `json:"id"`
	Hash        string              `json:"hash"`
	Transaction bitcoin.Transaction `json:"transaction"`
}
