package models

import (
	"errors"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/google/uuid"
)

type PayoutPayload struct {
	ID                    uuid.UUID
	Status                string
	Store_id              uuid.UUID
	Payment_id            uuid.UUID
	Type                  string
	Block_height_required int
	Transaction_hash      string
	Created_at            time.Time
	Action                string
}

func (p *PayoutPayload) Set_created_at() error {
	p.Created_at = time.Now()
	return nil
}

func (p *PayoutPayload) Set_id() error {
	p.ID = uuid.New()
	return nil
}

func (p *PayoutPayload) FromPayout(payout Payout) PayoutPayload {
	return PayoutPayload{
		ID:                    payout.ID,
		Status:                payout.Status,
		Store_id:              payout.Store_id,
		Payment_id:            payout.Payment_id,
		Type:                  payout.Type,
		Block_height_required: payout.Block_height_required,
		Transaction_hash:      payout.Transaction_hash,
		Created_at:            payout.Created_at,
		Action:                payout.Action,
	}
}

type Payout struct {
	ID                    uuid.UUID `json:"id"`
	Status                string    `json:"status"`
	Action                string    `json:"action"`
	Store_id              uuid.UUID `json:"store_id"`
	Payment_id            uuid.UUID `json:"payment_id"`
	Type                  string    `json:"crypto_type"`
	Block_height_required int       `json:"block_height_required"`
	Transaction_hash      string    `json:"transaction_hash"`
	Created_at            time.Time `json:"created_at"`
}

type InsertPayoutMessage struct {
	Payload Payout
}

func InsertPayout(e *actor.Engine, conn *actor.PID, d PayoutPayload) (Payout, error) {
	d.Set_created_at()
	d.Set_id()
	var resp = e.Request(conn, InsertPayoutMessage{
		Payload: d.ToPayout(),
	}, time.Millisecond*100)

	res, err := resp.Result()
	if err != nil {
		return Payout{}, errors.New("An error occured!")
	}
	myStruct, ok := res.(Payout)

	if !ok {
		return Payout{}, errors.New("An error occured!")
	}

	return myStruct, nil
}

type FindAllConfirmedPayoutMessage struct {
	Block_number int
	Crypto       string
}

func FindAllConfirmedPayout(e *actor.Engine, conn *actor.PID, block_number int, crypto string) ([]Payout, error) {
	var payouts []Payout

	var resp = e.Request(conn, FindAllConfirmedPayoutMessage{
		Block_number: block_number,
		Crypto:       crypto,
	}, time.Millisecond*100)

	res, err := resp.Result()

	if err != nil {
		return payouts, errors.New("An error occured!")
	}

	myStruct, ok := res.([]Payout)

	if !ok {
		return payouts, errors.New("An error occured!")
	}

	return myStruct, nil

}

type UpdatePayoutWithPaymentMessage struct {
	Payout_id       uuid.UUID
	Payout_payload  PayoutPayload
	Payment_payload PaymentPayload
}

func UpdatePayoutWithPayment(e *actor.Engine, conn *actor.PID, payout_id uuid.UUID, payout_payload PayoutPayload, payment_payload PaymentPayload) (Payout, error) {
	var resp = e.Request(conn, UpdatePayoutWithPaymentMessage{
		Payout_id:       payout_id,
		Payout_payload:  payout_payload,
		Payment_payload: payment_payload,
	}, time.Millisecond*100)
	res, err := resp.Result()
	if err != nil {
		return Payout{}, errors.New("An error occured!")
	}
	myStruct, ok := res.(Payout)

	if !ok {
		return Payout{}, errors.New("An error occured!")
	}

	return myStruct, nil
}

func (p *PayoutPayload) ToPayout() Payout {
	return Payout{
		ID:                    p.ID,
		Status:                p.Status,
		Store_id:              p.Store_id,
		Payment_id:            p.Payment_id,
		Type:                  p.Type,
		Block_height_required: p.Block_height_required,
		Transaction_hash:      p.Transaction_hash,
		Created_at:            p.Created_at,
		Action:                p.Action,
	}
}
