package dao

import (
	"strconv"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/server/util"
)

// TODO: blockheight_required
func CreatePayment(e *actor.Engine, conn *actor.PID, payload models.PaymentPayload) (*models.Payment, error) {
	payload.Set_id()
	price, _ := strconv.ParseFloat(payload.Price, 64)
	payload.TotalFee = price + payload.Fee + (0.05 * price)

	payload.Expires_at = time.Now().Add(30 * time.Minute)
	payment, err := models.InsertPayment(e, conn, payload)

	if err != nil {
		return nil, util.NewErrNotFound("payment creation error")
	}

	return &payment, nil
}
