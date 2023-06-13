package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/sijibomii/cryptopay/types"
	"github.com/sijibomii/cryptopay/types/bitcoin"
)

type StorePayload struct {
	ID                         uuid.UUID
	Name                       string
	Description                string
	Created_at                 time.Time
	Updated_at                 time.Time
	Owner_id                   uuid.UUID
	Private_key                types.PrivateKey
	Public_key                 types.PublicKey
	Btc_payout_addresses       []bitcoin.Address
	Btc_confirmations_required int
	Mnemonic                   string
	Hd_path                    string
	Deleted_at                 time.Time
}

func New() StorePayload {
	return StorePayload{}
}

func (sp *StorePayload) Set_created_at() error {
	sp.Created_at = time.Now()
	return nil
}

func (sp *StorePayload) Set_updated_at() error {
	sp.Updated_at = time.Now()
	return nil
}

func (sp *StorePayload) Set_deleted_at() error {
	sp.Name = ""
	sp.Description = ""
	sp.Deleted_at = time.Now()
	return nil
}

type Store struct {
	ID                         uuid.UUID
	Name                       string
	Description                string
	Created_at                 time.Time
	Updated_at                 time.Time
	Owner_id                   uuid.UUID
	Private_key                types.PrivateKey
	Public_key                 types.PublicKey
	Btc_payout_addresses       []bitcoin.Address
	Btc_confirmations_required int
	Mnemonic                   string
	Hd_path                    string
	Deleted_at                 time.Time
}
