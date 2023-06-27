package services

import (
	"strconv"

	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/server/util"
)

// currency api client should be in app state
func CreatePayment(appState *util.AppState, store models.Store, payload models.PaymentPayload) (*models.Payment, error) {
	payload.Set_created_at()
	payload.Set_updated_at()

	payload.Index = 1

	payload.Status = "pending"

	path := store.Hd_path

	path += "/"

	path += strconv.FormatInt(createdTime.Unix(), 10)
	path += "/"

	path += strconv.Itoa(createdTime.Nanosecond() / 1000)
}
