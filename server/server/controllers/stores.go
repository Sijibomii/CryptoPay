package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/server/services"
	"github.com/sijibomii/cryptopay/server/util"
)

type GetStoresResponse struct {
	Stores []models.Store `json:"stores"`
	Owner  string         `json:"owner_id"`
}

func GetStoresList(w http.ResponseWriter, r *http.Request, appState *util.AppState) {
	userContext := r.Context().Value("user").(models.User)

	queryValues := r.URL.Query()

	limit := queryValues.Get("limit")
	offset := queryValues.Get("offset")

	limitNum, limitErr := strconv.Atoi(limit)
	if limitErr != nil {
		limitNum = 15
	}
	offsetNum, offsetErr := strconv.Atoi(offset)

	if offsetErr != nil {
		offsetNum = 0
	}

	stores, err := services.FindStoresByOwnerId(appState, userContext.ID, offsetNum, limitNum)

	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	json, err := json.Marshal(GetStoresResponse{
		Stores: stores,
		Owner:  userContext.ID.String(),
	})
	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	util.JsonBytesResponse(w, http.StatusOK, json)
	return

}

func CreateStores(w http.ResponseWriter, r *http.Request, appState *util.AppState) {

}
