package dao

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/google/uuid"
	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/server/util"
)

func UpdatePayoutWithPayment(e *actor.Engine, conn *actor.PID, payout_id uuid.UUID, payout_payload models.PayoutPayload, payment_payload models.PaymentPayload) (*models.Payout, error) {
	payout, err := models.UpdatePayoutWithPayment(e, conn, payout_id, payout_payload, payment_payload)

	if err != nil {
		return nil, util.NewErrNotFound("payout failed to be created")
	}

	return &payout, nil
}
