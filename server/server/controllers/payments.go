package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/server/client"
	"github.com/sijibomii/cryptopay/server/services"
	"github.com/sijibomii/cryptopay/server/util"
	"github.com/sijibomii/cryptopay/types/currency"
)

type CreatePaymentParams struct {
	Price      string `json:"price"`
	Crypto     string `json:"crypto"`
	Fiat       string `json:"fiat"`
	Identifier string `json:"indentifier"`
}

type storeParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PaymentResponse struct {
	Payment models.Payment `json:"payment"`
	Token   string         `json:"token"`
	Store   storeParams    `json:"store"`
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

	payment, err := services.CreatePayment(appState, *store, payload)

	jwtPayload := util.JwtPayload{
		Client:     reqClient,
		Expires_at: payment.Expires_at,
	}

	key := os.Getenv("JWT_SECRET_KEY")
	token, err := jwtPayload.Encode(key)

	if err != nil {
		util.ErrorResponseFunc(w, r, util.NewErrUnauthorized("payment error (jwt)"))
		return
	}

	json, err := json.Marshal(PaymentResponse{
		Payment: *payment,
		Store: storeParams{
			Name:        store.Name,
			Description: store.Description,
		},
		Token: token,
	})
	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	util.JsonBytesResponse(w, http.StatusOK, json)
	return
}
