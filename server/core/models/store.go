package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/google/uuid"
	"github.com/sijibomii/cryptopay/types/bitcoin"
	"github.com/sijibomii/cryptopay/types/currency"
	"github.com/tyler-smith/go-bip32"
)

type StorePayload struct {
	ID                         uuid.UUID
	Name                       string
	Description                string
	Created_at                 time.Time
	Updated_at                 time.Time
	Owner_id                   uuid.UUID
	Private_key                bip32.Key
	Public_key                 bip32.Key
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
	ID          uuid.UUID
	Name        string
	Description string
	Created_at  time.Time
	Updated_at  time.Time
	Owner_id    uuid.UUID
	Private_key bip32.Key
	Public_key  bip32.Key
	// Btc_payout_addresses       string
	Btc_confirmations_required int
	Mnemonic                   string
	Hd_path                    string
	Deleted_at                 time.Time
}

func (sp *StorePayload) ToStore() Store {
	// addressStrings := make([]string, len(sp.Btc_payout_addresses))
	// for i, address := range sp.Btc_payout_addresses {
	// 	addressStrings[i] = string(address)
	// }
	// joinedAddresses := strings.Join(addressStrings, ",")

	return Store{
		ID:          sp.ID,
		Name:        sp.Name,
		Description: sp.Description,
		Created_at:  sp.Created_at,
		Updated_at:  sp.Updated_at,
		Owner_id:    sp.Owner_id,
		Private_key: sp.Private_key,
		Public_key:  sp.Public_key,
		// Btc_payout_addresses:       joinedAddresses,
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
		// if s.Btc_payout_addresses != "" && s.Btc_confirmations_required != 0 {
		// 	return true
		// }
		return false
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
	}, time.Millisecond*100)
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
	}, time.Millisecond*100)
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

func Find_Store_By_Owner_Id(e *actor.Engine, conn *actor.PID, id uuid.UUID, limit, offset int64) ([]Store, error) {
	var resp = e.Request(conn, FindStoreByOwnerMessage{
		OwnerID: id,
		Limit:   limit,
		Offset:  offset,
	}, time.Millisecond*100)
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
	}, time.Millisecond*100)
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
	}, time.Millisecond*100)
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
	}, time.Millisecond*100)
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
	}, time.Millisecond*100)
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
	}, time.Millisecond*100)
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

func (s *Store) Export() ([]byte, error) {

	// ss := strings.Split(s.Btc_payout_addresses, ",")
	// addressStrings := make([]bitcoin.Address, len(ss))
	// for i, address := range s.Btc_payout_addresses {
	// 	addressStrings[i] = bitcoin.Address(address)
	// }

	data := struct {
		ID                         uuid.UUID         `json:"id"`
		Description                string            `json:"description"`
		Name                       string            `json:"name"`
		Btc_payout_addresses       []bitcoin.Address `json:"btc_payout_addresses"`
		Btc_confirmations_required int               `json:"btc_confirmations_required"`
		Public_key                 bip32.Key         `json:"public_key"`
		Can_accept_btc             bool              `json:"can_accept_btc"`
		CreatedAt                  time.Time         `json:"created_at"`
		UpdatedAt                  time.Time         `json:"updated_at"`
	}{
		ID:          s.ID,
		Description: s.Description,
		Name:        s.Name,
		// Btc_payout_addresses:       addressStrings,
		Btc_confirmations_required: s.Btc_confirmations_required,
		Public_key:                 s.Public_key,
		Can_accept_btc:             false, //s.Can_accept(currency.Btc),
	}

	return json.Marshal(data)
}
