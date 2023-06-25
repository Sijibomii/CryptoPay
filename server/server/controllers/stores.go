package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sijibomii/cryptopay/core/models"
	"github.com/sijibomii/cryptopay/server/services"
	"github.com/sijibomii/cryptopay/server/util"
)

type GetStoresResponse struct {
	Stores []models.Store `json:"stores"`
	Owner  string         `json:"owner_id"`
}

type CreateStoreParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateStoreReponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateStoreParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetStoresList(w http.ResponseWriter, r *http.Request, appState *util.AppState) {
	userContext := r.Context().Value("user").(*models.User)

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
	userContext := r.Context().Value("user").(*models.User)

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	defer r.Body.Close()

	var createStoreData CreateStoreParams
	err = json.Unmarshal(requestBody, &createStoreData)

	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	store, err := services.CreateStore(appState, userContext.ID, createStoreData.Name, createStoreData.Description)

	json, err := store.Export()

	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	util.JsonBytesResponse(w, http.StatusOK, json)
	return
}

func GetStoresById(w http.ResponseWriter, r *http.Request, appState *util.AppState) {
	userContext := r.Context().Value("user").(models.User)

	vars := mux.Vars(r)
	id := vars["id"]

	storeId, tokenErr := uuid.Parse(id)

	if tokenErr != nil {
		util.ErrorResponseFunc(w, r, tokenErr)
		return
	}

	store, err := services.FindStoreById(appState, storeId)

	if store.Owner_id != userContext.ID {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	json, err := store.Export()

	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	util.JsonBytesResponse(w, http.StatusOK, json)
	return
}

func UpdateStoresById(w http.ResponseWriter, r *http.Request, appState *util.AppState) {
	userContext := r.Context().Value("user").(*models.User)

	vars := mux.Vars(r)
	id := vars["id"]

	storeId, tokenErr := uuid.Parse(id)

	if tokenErr != nil {
		util.ErrorResponseFunc(w, r, tokenErr)
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	defer r.Body.Close()

	// what happens if i only update btc?
	var updateStoreData UpdateStoreParams
	err = json.Unmarshal(requestBody, &updateStoreData)

	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	store, err := services.FindStoreById(appState, storeId)

	if store.Owner_id != userContext.ID {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// update store
	store, err = services.UpdateStoresById(appState, storeId, updateStoreData.Name, updateStoreData.Description)

	json, err := store.Export()

	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	util.JsonBytesResponse(w, http.StatusOK, json)
	return
}
