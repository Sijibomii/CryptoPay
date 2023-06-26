package services

import (
	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/server/util"
)

// currency api client should be in app state
func CreatePayment(appState *util.AppState, store models.Store, payload models.PaymentPayload)
