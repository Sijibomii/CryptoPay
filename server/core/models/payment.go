package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/google/uuid"
)

type PaymentPayload struct {
	ID                     uuid.UUID
	Status                 string
	Store_id               uuid.UUID
	Index                  int
	TotalFee               float64
	Created_by             uuid.UUID
	Created_at             time.Time
	Updated_at             time.Time
	Expires_at             time.Time
	Paid_at                time.Time
	Amount_paid            string
	Transaction_hash       string
	Fiat                   string
	Price                  string
	Crypto                 string
	Address                string
	Charge                 string
	Confirmations_required int
	Block_height_required  int
	Btc_network            string
	Identifier             string
	Fee                    float64
}

func (p *PaymentPayload) Set_created_at() error {
	p.Created_at = time.Now()
	return nil
}

func (p *PaymentPayload) Set_paid_at() error {
	p.Paid_at = time.Now()
	return nil
}

func (p *PaymentPayload) Set_updated_at() error {
	p.Updated_at = time.Now()
	return nil
}

func (p *PaymentPayload) Set_id() error {
	p.ID = uuid.New()
	return nil
}

// i used str for double bc golang doesn't have an impl. I'll use an external library to convert
type Payment struct {
	ID                     uuid.UUID `json:"id"`
	Status                 string    `json:"payment_status"`
	Store_id               uuid.UUID `json:"store_id"`
	Index                  int       `json:"index"`
	Created_by             uuid.UUID `json:"created_by"`
	Created_at             time.Time `json:"created_at"`
	Updated_at             time.Time `json:"updated_at"`
	Expires_at             time.Time `json:"expires_at"`
	Paid_at                time.Time `json:"paid_at"`
	Amount_paid            string    `json:"amount_paid"`
	Transaction_hash       string    `json:"transaction_hash"`
	Fiat                   string    `json:"fiat"`
	Price                  string    `json:"price"`
	Crypto                 string    `json:"crypto"`
	Address                string    `json:"address"`
	Charge                 string    `json:"charge"`
	Confirmations_required int       `json:"confirmations_required"`
	Block_height_required  int       `json:"block_height_required"`
	Btc_network            string    `json:"btc_network"`
	Identifier             string    `json:"identifier"`
	TotalFee               float64   `json:"total"`
	Fee                    float64   `json:"miners_fee"`
}

func (p *PaymentPayload) FromPayment(payment Payment) PaymentPayload {
	return PaymentPayload{
		ID:                     payment.ID,
		Status:                 payment.Status,
		Store_id:               payment.Store_id,
		Index:                  payment.Index,
		Created_by:             payment.Created_by,
		Created_at:             payment.Created_at,
		Updated_at:             payment.Updated_at,
		Expires_at:             payment.Expires_at,
		Paid_at:                payment.Paid_at,
		Amount_paid:            payment.Amount_paid,
		Transaction_hash:       payment.Transaction_hash,
		Fiat:                   payment.Fiat,
		Price:                  payment.Price,
		Crypto:                 payment.Crypto,
		Address:                payment.Address,
		Charge:                 payment.Charge,
		Confirmations_required: payment.Confirmations_required,
		Block_height_required:  payment.Block_height_required,
		Btc_network:            payment.Btc_network,
		Identifier:             payment.Identifier,
		Fee:                    payment.Fee,
	}
}

func (p *PaymentPayload) ToPayment() Payment {
	return Payment{
		ID:                     p.ID,
		Status:                 p.Status,
		Store_id:               p.Store_id,
		Index:                  p.Index,
		Created_by:             p.Created_by,
		Created_at:             p.Created_at,
		Updated_at:             p.Updated_at,
		Expires_at:             p.Expires_at,
		Paid_at:                p.Paid_at,
		Amount_paid:            p.Amount_paid,
		Transaction_hash:       p.Transaction_hash,
		Fiat:                   p.Fiat,
		Price:                  p.Price,
		Crypto:                 p.Crypto,
		Address:                p.Address,
		Charge:                 p.Charge,
		Confirmations_required: p.Confirmations_required,
		Block_height_required:  p.Block_height_required,
		Btc_network:            p.Btc_network,
		Identifier:             p.Identifier,
		Fee:                    p.Fee,
	}
}

type InsertPaymentMessage struct {
	Payload Payment
}

func InsertPayment(e *actor.Engine, conn *actor.PID, d PaymentPayload) (Payment, error) {
	d.Set_created_at()
	var resp = e.Request(conn, InsertPaymentMessage{
		Payload: d.ToPayment(),
	}, time.Millisecond*100)

	res, err := resp.Result()
	if err != nil {
		return Payment{}, errors.New("An error occured!")
	}
	myStruct, ok := res.(Payment)

	if !ok {
		return Payment{}, errors.New("An error occured!")
	}

	return myStruct, nil
}

type FindAllPendingPaymentByAddressesMessage struct {
	Address []string
	// btc
	Crypto string
}

func FindAllPendingPaymentsByAddresses(e *actor.Engine, conn *actor.PID, address []string, crypto string) ([]Payment, error) {
	var resp = e.Request(conn, FindAllPendingPaymentByAddressesMessage{
		Address: address,
		Crypto:  crypto,
	}, time.Millisecond*100)

	res, err := resp.Result()
	var cts []Payment
	if err != nil {
		return cts, errors.New("An error occured!")
	}
	myStruct, ok := res.([]Payment)

	if !ok {
		return cts, errors.New("An error occured!")
	}

	return myStruct, nil
}

type UpdatePaymentMessage struct {
	Payload Payment
	Id      uuid.UUID
}

func UpdatePayment(e *actor.Engine, conn *actor.PID, id uuid.UUID, d PaymentPayload) (Payment, error) {
	var resp = e.Request(conn, UpdatePaymentMessage{
		Payload: d.ToPayment(),
		Id:      id,
	}, time.Millisecond*100)

	res, err := resp.Result()

	if err != nil {
		return Payment{}, errors.New("An error occured!")
	}
	myStruct, ok := res.(Payment)

	if !ok {
		return Payment{}, errors.New("An error occured!")
	}

	return myStruct, nil
}

func (p *Payment) Export() ([]byte, error) {
	data := struct {
		ID                     uuid.UUID `json:"id"`
		Status                 string    `json:"payment_status"`
		Store_id               uuid.UUID `json:"store_id"`
		Index                  int       `json:"index"`
		Created_by             uuid.UUID `json:"created_by"`
		Created_at             time.Time `json:"created_at"`
		Updated_at             time.Time `json:"updated_at"`
		Expires_at             time.Time `json:"expires_at"`
		Paid_at                time.Time `json:"paid_at"`
		Amount_paid            string    `json:"amount_paid"`
		Transaction_hash       string    `json:"transaction_hash"`
		Fiat                   string    `json:"fiat"`
		Price                  string    `json:"price"`
		Crypto                 string    `json:"crypto"`
		Address                string    `json:"address"`
		Charge                 string    `json:"charge"`
		Confirmations_required int       `json:"confirmations_required"`
		Block_height_required  int       `json:"block_height_required"`
		Btc_network            string    `json:"btc_network"`
		Identifier             string    `json:"identifier"`
		TotalFee               float64   `json:"total"`
		Fee                    float64   `json:"miners_fee"`
	}{
		ID:                     p.ID,
		Status:                 p.Status,
		Store_id:               p.Store_id,
		Index:                  p.Index,
		Created_by:             p.Created_by,
		Created_at:             p.Created_at,
		Updated_at:             p.Updated_at,
		Expires_at:             p.Expires_at,
		Paid_at:                p.Paid_at,
		Amount_paid:            p.Amount_paid,
		Transaction_hash:       p.Transaction_hash,
		Fiat:                   p.Fiat,
		Price:                  p.Price,
		Crypto:                 p.Crypto,
		Address:                p.Address,
		Charge:                 p.Charge,
		Confirmations_required: p.Confirmations_required,
		Block_height_required:  p.Block_height_required,
		Btc_network:            p.Btc_network,
		TotalFee:               p.TotalFee,
		Identifier:             p.Identifier,
		Fee:                    p.Fee,
	}

	return json.Marshal(data)
}
