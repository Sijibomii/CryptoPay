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

type RegisterParams struct {
	email    string
	password string
}

func LoginHandler(w http.ResponseWriter, r *http.Request, appState *util.AppState) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	var loginData LoginParams
	err = json.Unmarshal(requestBody, &loginData)
	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	token, err := services.Login(appState, loginData.email, loginData.password)
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

func RegisterHandler(w http.ResponseWriter, r *http.Request, appState *util.AppState) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	var registerData RegisterParams
	err = json.Unmarshal(requestBody, &registerData)
	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}
	token, err := services.Login(appState, registerData.email, registerData.password)
}
