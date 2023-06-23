package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/server/services"
	"github.com/sijibomii/cryptopay/server/util"
)

type ClientTokensResponse struct {
	Tokens   []models.ClientToken `json:"tokens"`
	Store_id uuid.UUID            `json:"store_id"`
	Offset   int                  `json:"offset"`
	Limit    int                  `json:"limit"`
}

func GetAllClientTokensHandler(w http.ResponseWriter, r *http.Request, appState *util.AppState) {
	userContext := r.Context().Value("user").(models.User)

	queryValues := r.URL.Query()

	limit := queryValues.Get("limit")
	offset := queryValues.Get("offset")
	store_id := queryValues.Get("store")

	limitNum, limitErr := strconv.Atoi(limit)
	if limitErr != nil {
		limitNum = 15
	}
	offsetNum, offsetErr := strconv.Atoi(offset)

	if offsetErr != nil {
		offsetNum = 0
	}

	id, tokenErr := uuid.Parse(store_id)

	if tokenErr != nil {
		util.ErrorResponseFunc(w, r, tokenErr)
		return
	}

	store, err := services.FindStoreById(appState, id)

	if err != nil {
		util.ErrorResponseFunc(w, r, tokenErr)
		return
	}

	if store.Owner_id != userContext.ID {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokens, tokensErr := services.GetAllClientTokensByStoreId(appState, store.ID, offsetNum, limitNum)

	if tokenErr != nil {
		util.ErrorResponseFunc(w, r, tokensErr)
		return
	}

	json, err := json.Marshal(ClientTokensResponse{
		Tokens:   tokens,
		Store_id: store.ID,
		Offset:   offsetNum,
		Limit:    limitNum,
	})
	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	util.JsonBytesResponse(w, http.StatusOK, json)
	return
}

func CreateClientTokensHandler(w http.ResponseWriter, r *http.Request, appState *util.AppState) {

}
