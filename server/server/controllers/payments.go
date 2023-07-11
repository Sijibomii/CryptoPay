package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/server/client"
	"github.com/sijibomii/cryptopay/server/dao"
	"github.com/sijibomii/cryptopay/server/services"
	"github.com/sijibomii/cryptopay/server/util"
	"github.com/sijibomii/cryptopay/types/currency"
)

type CreatePaymentParams struct {
	Price      int    `json:"price"`
	Crypto     string `json:"crypto"`
	Fiat       string `json:"fiat"`
	Identifier string `json:"identifier"`
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

	store, err := services.FindStoreById(appState, clientTokenContext.Store_id)

	var crypto currency.Crypto
	err = crypto.Scan([]byte(createPaymentData.Crypto))
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
		Price:      strconv.Itoa(createPaymentData.Price),
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
	fmt.Printf("jwt secret key: %s", key)

	token, err := jwtPayload.Encode(appState.PrivateKey)

	if err != nil {
		fmt.Print("\n error", err.Error())
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

type PaymentStatusResponse struct {
	Status                  string `json:"status"`
	Confirmations_required  int    `json:"confirmations_required"`
	Remaining_confirmations int    `json:"remaining_confirmations"`
}

// payment status a token was sent back to the user, that will be decoded and be sent now instead of api key
func GetPaymentStatus(w http.ResponseWriter, r *http.Request, appState *util.AppState) {

	jwtPayload := r.Context().Value("Payload").(*util.JwtPayload)

	vars := mux.Vars(r)
	id := vars["id"]

	paymentId, tokenErr := uuid.Parse(id)

	if tokenErr != nil {
		util.ErrorResponseFunc(w, r, tokenErr)
		return
	}

	payment, err := services.GetPaymentById(appState, paymentId)

	// validate payment owner
	if payment.Created_by != jwtPayload.Client.ID || err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	switch payment.Crypto {
	case "btc":
		//get the current payment status
		status, err := dao.FindBtcBlockChainStatusByNetwork(appState.Engine, appState.Postgres, payment.Btc_network)

		remaining_conf := payment.Confirmations_required - status.Block_Height

		if payment.Status == "paid" {
			remaining_conf = 0
		}

		block_height_required := payment.Block_height_required

		if block_height_required < status.Block_Height {
			remaining_conf = 0
		} else {
			remaining_conf = block_height_required - status.Block_Height
		}

		json, err := json.Marshal(PaymentStatusResponse{
			Remaining_confirmations: remaining_conf,
			Confirmations_required:  payment.Confirmations_required,
			Status:                  payment.Status,
		})
		if err != nil {
			util.ErrorResponseFunc(w, r, err)
			return
		}

		util.JsonBytesResponse(w, http.StatusOK, json)
		return

	}

}
