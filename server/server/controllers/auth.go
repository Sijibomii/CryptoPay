package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sijibomii/cryptopay/server/services"
	"github.com/sijibomii/cryptopay/server/util"
)

type LoginParams struct {
	email    string
	password string
}

type LoginResponse struct {
	// Session token
	// required: true
	Token string `json:"token"`
}

// TODO: create error response in utils, stringResponse, jsonStringResponse, jsonBytesResponse, setResponseHeader

func LoginHandler(w http.ResponseWriter, r *http.Request, appState *util.AppState) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		// error response
		util.ErrorResponseFunc(w, r, err)
		return
	}

	var loginData LoginParams
	err = json.Unmarshal(requestBody, &loginData)
	if err != nil {
		// error response
		util.ErrorResponseFunc(w, r, err)
		return
	}

	token, err := services.Login(loginData.email, loginData.password)
	if err != nil {
		util.ErrorResponseFunc(w, r, util.NewErrUnauthorized("incorrect login"))
		return
	}
	json, err := json.Marshal(LoginResponse{Token: token})
	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	util.JsonBytesResponse(w, http.StatusOK, json)
	return
}
