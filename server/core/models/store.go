package models

import (
	"encoding/json"
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

func (s *Store) Can_accept(crypto any) bool {
	return false
}

func InsertStore(UserPayload, d UserPayload) error {
	return nil
}

func UpdateStore(UserPayload, d UserPayload) error {
	return nil
}

func Find_By_Owner(UserPayload, d UserPayload) error {
	return nil
}

func Find_By_Id(UserPayload, d UserPayload) error {
	return nil
}

func Find_By_Id_With_Deleted(UserPayload, d UserPayload) error {
	return nil
}
func Soft_Delete() error {
	return nil
}

func (s *Store) export() ([]byte, error) {
	data := struct {
		ID                         uuid.UUID         `json:"id"`
		Description                string            `json:"description"`
		Name                       string            `json:"name"`
		Btc_payout_addresses       []bitcoin.Address `json:"btc_payout_addresses"`
		Btc_confirmations_required int               `json:"btc_confirmations_required"`
		Public_key                 types.PublicKey   `json:"public_key"`
		Can_accept_btc             bool              `json:"can_accept_btc"`
		CreatedAt                  time.Time         `json:"created_at"`
		UpdatedAt                  time.Time         `json:"updated_at"`
	}{
		ID:                         s.ID,
		Description:                s.Description,
		Name:                       s.Name,
		Btc_payout_addresses:       s.Btc_payout_addresses,
		Btc_confirmations_required: s.Btc_confirmations_required,
		Public_key:                 s.Public_key,
		// change to crypto type
		Can_accept_btc: s.Can_accept("btc"),
		CreatedAt:      s.Created_at,
		UpdatedAt:      s.Updated_at,
	}

	return json.Marshal(data)
}
