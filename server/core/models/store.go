package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/google/uuid"
	"github.com/sijibomii/cryptopay/types"
	"github.com/sijibomii/cryptopay/types/bitcoin"
	"github.com/sijibomii/cryptopay/types/currency"
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
	Btc_payout_addresses       []bitcoin.Address `gorm:"type:text[]"`
	Btc_confirmations_required int
	Mnemonic                   string
	Hd_path                    string
	Deleted_at                 time.Time
}

func (sp *StorePayload) ToStore() Store {
	return Store{
		ID:                         sp.ID,
		Name:                       sp.Name,
		Description:                sp.Description,
		Created_at:                 sp.Created_at,
		Updated_at:                 sp.Updated_at,
		Owner_id:                   sp.Owner_id,
		Private_key:                sp.Private_key,
		Public_key:                 sp.Public_key,
		Btc_payout_addresses:       sp.Btc_payout_addresses,
		Btc_confirmations_required: sp.Btc_confirmations_required,
		Mnemonic:                   sp.Mnemonic,
		Hd_path:                    sp.Hd_path,
		Deleted_at:                 sp.Deleted_at,
	}
}

type InsertStoreMessage struct {
	Payload Store
}

type UpdateStoreMessage struct {
	Payload Store
	Id      uuid.UUID
}

type FindStoreByIdMessage struct {
	Id uuid.UUID
}

type FindStoreByOwnerMessage struct {
	OwnerID uuid.UUID
	Limit   int64
	Offset  int64
}

type FindStoreByIdWithDeletedMessage struct {
	Id uuid.UUID
}

type DeleteStoreMessage struct {
	Id uuid.UUID
}

type SoftDeleteStoreMessage struct {
	Id uuid.UUID
}

type SoftDeleteStoreByOwnerIDMessage struct {
	OwnerID uuid.UUID
}

func (s *Store) Can_accept(crypto currency.Crypto) bool {
	switch crypto {
	case currency.Btc:
		if s.Btc_payout_addresses != nil && s.Btc_confirmations_required != 0 {
			return true
		}
	case currency.Eth:
		fmt.Println("Ethereum (ETH)")
	default:
		return false
	}
	return false
}

// confirm the return type before compl code check users too
func InsertStore(e *actor.Engine, conn *actor.PID, d StorePayload) (Store, error) {
	d.Set_created_at()
	var resp = e.Request(conn, InsertStoreMessage{
		Payload: d.ToStore(),
	}, 500)
	res, err := resp.Result()
	if err != nil {
		return Store{}, errors.New("An error occured!")
	}
	myStruct, ok := res.(Store)

	if !ok {
		return Store{}, errors.New("An error occured!")
	}

	return myStruct, nil
}

func UpdateStore(e *actor.Engine, conn *actor.PID, id uuid.UUID, d StorePayload) (Store, error) {
	var resp = e.Request(conn, UpdateStoreMessage{
		Payload: d.ToStore(),
		Id:      id,
	}, 500)
	res, err := resp.Result()
	if err != nil {
		return Store{}, errors.New("An error occured!")
	}
	myStruct, ok := res.(Store)

	if !ok {
		return Store{}, errors.New("An error occured!")
	}

	return myStruct, nil
}

func Find_By_Owner(e *actor.Engine, conn *actor.PID, id uuid.UUID, limit, offset int64) ([]Store, error) {
	var resp = e.Request(conn, FindStoreByOwnerMessage{
		OwnerID: id,
		Limit:   limit,
		Offset:  offset,
	}, 500)
	res, err := resp.Result()
	var stores []Store
	if err != nil {
		return stores, errors.New("An error occured!")
	}
	myStruct, ok := res.([]Store)

	if !ok {
		return stores, errors.New("An error occured!")
	}

	return myStruct, nil
}

func Find_Store_By_Id(e *actor.Engine, conn *actor.PID, id uuid.UUID) (Store, error) {
	var resp = e.Request(conn, FindStoreByIdMessage{
		Id: id,
	}, 500)
	res, err := resp.Result()
	if err != nil {
		return Store{}, errors.New("An error occured!")
	}
	myStruct, ok := res.(Store)

	if !ok {
		return Store{}, errors.New("An error occured!")
	}
	return myStruct, nil
}

func Find_By_Id_With_Deleted(e *actor.Engine, conn *actor.PID, id uuid.UUID) (Store, error) {
	var resp = e.Request(conn, FindStoreByIdWithDeletedMessage{
		Id: id,
	}, 500)
	res, err := resp.Result()
	if err != nil {
		return Store{}, errors.New("An error occured!")
	}
	myStruct, ok := res.(Store)

	if !ok {
		return Store{}, errors.New("An error occured!")
	}
	return myStruct, nil
}

func Soft_Delete(e *actor.Engine, conn *actor.PID, id uuid.UUID) (bool, error) {
	var resp = e.Request(conn, SoftDeleteStoreMessage{
		Id: id,
	}, 500)
	res, err := resp.Result()
	if err != nil {
		return false, errors.New("An error occured!")
	}
	_, ok := res.(bool)

	if !ok {
		return false, errors.New("An error occured!")
	}
	return true, nil
}

func Delete(e *actor.Engine, conn *actor.PID, id uuid.UUID) (int64, error) {
	var resp = e.Request(conn, DeleteStoreMessage{
		Id: id,
	}, 500)
	res, err := resp.Result()
	if err != nil {
		return 0, errors.New("An error occured!")
	}
	num, ok := res.(int64)

	if !ok {
		return 0, errors.New("An error occured!")
	}
	return num, nil
}

func Soft_Delete_Store_By_OwnerID(e *actor.Engine, conn *actor.PID, Owner_id uuid.UUID) (bool, error) {
	var resp = e.Request(conn, SoftDeleteStoreByOwnerIDMessage{
		OwnerID: Owner_id,
	}, 500)
	res, err := resp.Result()
	if err != nil {
		return false, errors.New("An error occured!")
	}
	_, ok := res.(bool)

	if !ok {
		return false, errors.New("An error occured!")
	}
	return true, nil
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
		Can_accept_btc:             s.Can_accept(currency.Btc),
		CreatedAt:                  s.Created_at,
		UpdatedAt:                  s.Updated_at,
	}

	return json.Marshal(data)
}
