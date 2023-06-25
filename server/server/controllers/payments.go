package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/server/client"
	"github.com/sijibomii/cryptopay/server/services"
	"github.com/sijibomii/cryptopay/server/util"
	"github.com/sijibomii/cryptopay/types/bitcoin"
	"github.com/sijibomii/cryptopay/types/currency"
)

type CreatePaymentParams struct {
	Price      string `json:"price"`
	Crypto     string `json:"crypto"`
	Fiat       string `json:"fiat"`
	Identifier string `json:"indentifier"`
}

// api key comes in auth header
func CreatePayment(w http.ResponseWriter, r *http.Request, appState *util.AppState) {
	clientTokenContext := r.Context().Value("Ctoken").(*models.ClientToken)

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	defer r.Body.Close()

	var createPaymentData CreatePaymentParams
	err = json.Unmarshal(requestBody, &createPaymentData)

	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	store, err := services.FindStoreById(appState, clientTokenContext.ID)

	var crypto currency.Crypto
	err = crypto.Scan(createPaymentData.Crypto)
	if err != nil {
		// return response
		http.Error(w, "crypto field is invalid", http.StatusBadRequest)
		return
	}

	if !store.Can_accept(crypto) {
		http.Error(w, "crypto not accepted by this store", http.StatusBadRequest)
		return
	}

	// create a client struct
	reqClient := client.NewClient(*clientTokenContext)

	payload := models.PaymentPayload{
		Store_id:   reqClient.Store_id,
		Created_by: reqClient.ID,
		Fiat:       createPaymentData.Fiat,
		Price:      createPaymentData.Price,
		Crypto:     createPaymentData.Crypto,
	}

	if identifier := createPaymentData.Identifier; identifier != "" {
		if len(identifier) > 100 {
			http.Error(w, "identifier is too long. Maximum of 100", http.StatusBadRequest)
			return
		}
		payload.Identifier = identifier
	} else {
		http.Error(w, "missing identifier", http.StatusBadRequest)
		return
	}

	var min_charge string

	switch crypto {
	case currency.Btc:
		payload.Confirmations_required = store.Btc_confirmations_required
		payload.Btc_network = bitcoin.Test.String()
		min_charge = "0.01"

	}

}
