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

	path += strconv.FormatInt(payload.Created_at.Unix(), 10)
	path += "/"

	path += strconv.Itoa(payload.Created_at.Nanosecond() / 1000)

	// rate, err := util.GetRate(appState.Engine, appState.CoinClient, payload.Fiat, payload.Crypto)
	// price, err := strconv.ParseFloat(payload.Price, 64)
	// if err != nil {
	// 	// Handle the error if the string cannot be parsed
	// 	fmt.Println("Error converting string to float64:", err)
	// 	// return
	// }

	// charge := rate * float64(price)

	// charges
	return &models.Payment{}, nil
}
