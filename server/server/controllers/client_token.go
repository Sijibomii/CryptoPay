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

type ClientTokensResponse struct {
	Tokens   []models.ClientToken `json:"tokens"`
	Store_id uuid.UUID            `json:"store_id"`
	Offset   int                  `json:"offset"`
	Limit    int                  `json:"limit"`
}

type CreateTokenParams struct {
	Name     string `json:"name"`
	Store_id string `json:"store_id"`
	Domain   string `json:"domain"`
}

type CreateTokenResponse struct {
	Name     string `json:"name"`
	Store_id string `json:"store_id"`
	Domain   string `json:"domain"`
	Token    string `json:"token"`
}

type GetClientTokenById struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Store_id string `json:"store_id"`
	Domain   string `json:"domain"`
	Token    string `json:"token"`
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

func GetClientTokenByIdHandler(w http.ResponseWriter, r *http.Request, appState *util.AppState) {
	userContext := r.Context().Value("user").(models.User)

	vars := mux.Vars(r)
	id := vars["id"]

	clientTokenId, tokenErr := uuid.Parse(id)

	if tokenErr != nil {
		util.ErrorResponseFunc(w, r, tokenErr)
		return
	}

	token, err := services.GetClientTokenById(appState, clientTokenId)

	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	store, err := services.FindStoreById(appState, token.Store_id)

	if err != nil {
		util.ErrorResponseFunc(w, r, tokenErr)
		return
	}

	if store.Owner_id != userContext.ID {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	json, err := json.Marshal(GetClientTokenById{
		ID:       token.ID.String(),
		Name:     token.Name,
		Domain:   token.Domain,
		Store_id: token.Store_id.String(),
		Token:    token.Token.String(),
	})
	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	util.JsonBytesResponse(w, http.StatusOK, json)
	return
}

func CreateClientTokensHandler(w http.ResponseWriter, r *http.Request, appState *util.AppState) {
	userContext := r.Context().Value("user").(*models.User)

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	defer r.Body.Close()

	var createTokenData CreateTokenParams
	err = json.Unmarshal(requestBody, &createTokenData)

	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	id, tokenErr := uuid.Parse(createTokenData.Store_id)

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

	token, tokensErr := services.CreateClientToken(appState, store.ID, createTokenData.Name, createTokenData.Domain)

	if tokenErr != nil {
		util.ErrorResponseFunc(w, r, tokensErr)
		return
	}

	json, err := json.Marshal(CreateTokenResponse{
		Token:    token.Token.String(),
		Name:     token.Name,
		Domain:   token.Domain,
		Store_id: token.Store_id.String(),
	})
	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	util.JsonBytesResponse(w, http.StatusOK, json)
	return

}
