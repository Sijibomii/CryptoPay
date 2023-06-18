package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/sijibomii/cryptopay/server/services"
	"github.com/sijibomii/cryptopay/server/util"
)

type LoginParams struct {
	Email    string
	Password string
}

type LoginResponse struct {
	// Session token
	// required: true
	Token string `json:"token"`
}

type RegisterParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct{}

type ActivationResponse struct{}

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

	token, err := services.Login(appState, loginData.Email, loginData.Password)
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

	defer r.Body.Close()

	var registerData RegisterParams
	err = json.Unmarshal(requestBody, &registerData)
	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}
	fmt.Printf("EMAIL IS %s ################ \n", registerData.Email)
	registerErr := services.Register(appState, registerData.Email, registerData.Password)

	if registerErr != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}
	json, err := json.Marshal(RegisterResponse{})

	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	util.JsonBytesResponse(w, http.StatusOK, json)
	return
}

func Activation(w http.ResponseWriter, r *http.Request, appState *util.AppState) {
	queryParams, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Fatal(err)
	}

	token := queryParams.Get("activation")

	err = services.Activate(appState, token)

	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}
	json, err := json.Marshal(ActivationResponse{})

	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	util.JsonBytesResponse(w, http.StatusOK, json)
	return
}
