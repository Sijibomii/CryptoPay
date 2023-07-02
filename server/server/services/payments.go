package services

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sijibomii/cryptopay/blockchain_client/bitcoin"
	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/server/dao"
	"github.com/sijibomii/cryptopay/server/util"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

// currency api client should be in app state

func CreatePayment(appState *util.AppState, store models.Store, payload models.PaymentPayload) (*models.Payment, error) {
	payload.Set_created_at()
	payload.Set_updated_at()

	payload.Confirmations_required = store.Btc_confirmations_required
	payload.Index = 1

	payload.Btc_network = "TEST"

	payload.Status = "pending"

	path := store.Hd_path

	path += "/"

	path += strconv.FormatInt(payload.Created_at.Unix(), 10)
	path += "/"

	path += strconv.Itoa(payload.Created_at.Nanosecond() / 1000)

	//create the address for this trans
	stores_mnemonic := store.Mnemonic
	seed := bip39.NewSeed(stores_mnemonic, "")

	masterKey, _ := bip32.NewMasterKey(seed)

	childKey, _ := util.NewChildKeyFromString(masterKey, path)

	payload.Address = childKey.PublicKey().String()

	rate, err := util.GetRate(appState.Engine, appState.CoinClient, payload.Fiat, payload.Crypto)

	price, err := strconv.ParseFloat(payload.Price, 64)

	if err != nil {
		// Handle the error if the string cannot be parsed
		fmt.Println("Error converting string to float64:", err)
		return &models.Payment{}, errors.Wrap(err, "internal error")
	}
	fee, err := bitcoin.GetFeeEstimates(appState.Engine, appState.BtcClient)

	charge := (rate * float64(price)) + fee.OneHourFee

	payload.Charge = strconv.FormatFloat(charge, 'f', -1, 64)
	payload.Fee = fee.OneHourFee
	payment, err := dao.CreatePayment(appState.Engine, appState.Postgres, payload)

	return payment, nil
}

func GetPaymentById(appState *util.AppState, payment_id uuid.UUID) (*models.Payment, error) {
	var payment *models.Payment
	var err error

	payment, err = dao.GetPaymentById(appState.Engine, appState.Postgres, payment_id)

	if err != nil {
		return nil, errors.Wrap(err, "error geting payment")
	}

	return payment, nil
}
