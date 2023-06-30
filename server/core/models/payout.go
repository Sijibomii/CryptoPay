package models

import (
	"time"

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
}

func (p *PayoutPayload) Set_created_at() error {
	p.Created_at = time.Now()
	return nil
}

func (p *PayoutPayload) Set_id() error {
	p.ID = uuid.New()
	return nil
}

type Payout struct {
	ID                    uuid.UUID `json:"id"`
	Status                string    `json:"status"`
	Store_id              uuid.UUID `json:"store_id"`
	Payment_id            uuid.UUID `json:"payment_id"`
	Type                  string    `json:"crypto_type"`
	Block_height_required int       `json:"block_height_required"`
	Transaction_hash      string    `json:"transaction_hash"`
	Created_at            time.Time `json:"created_at"`
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
	}
}
