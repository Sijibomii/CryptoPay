package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/google/uuid"
	"github.com/sijibomii/cryptopay/server/services"
	"github.com/sijibomii/cryptopay/server/util"
)

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

type RegisterResponse struct {
	ID          uuid.UUID `json:"id"`
	Email       string    `json:"email"`
	Is_verified bool      `json:"isVerified"`
}

type ActivationResponse struct{}

type ResetPasswordParams struct {
	Email string `json:"email"`
}

type ResetPasswordResponse struct {
	Success bool `json:"success"`
}

type ChangePasswordParams struct {
	Password string `json:"password"`
	Token    string `json:"token"`
}

func ResetPasswordHandler(w http.ResponseWriter, r *http.Request, appState *util.AppState) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	defer r.Body.Close()
	var resetData ResetPasswordParams
	err = json.Unmarshal(requestBody, &resetData)
	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}
	_, err = services.ResetPassword(appState, resetData.Email)

	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	json, err := json.Marshal(ResetPasswordResponse{Success: true})
	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	util.JsonBytesResponse(w, http.StatusOK, json)
	return
}

func ChangePasswordHandler(w http.ResponseWriter, r *http.Request, appState *util.AppState) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	defer r.Body.Close()
}

func LoginHandler(w http.ResponseWriter, r *http.Request, appState *util.AppState) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}

	defer r.Body.Close()

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

	registeredUser, registerErr := services.Register(appState, registerData.Email, registerData.Password)

	if registerErr != nil {
		util.ErrorResponseFunc(w, r, err)
		return
	}
	json, err := json.Marshal(RegisterResponse{
		ID:          registeredUser.ID,
		Email:       registeredUser.Email,
		Is_verified: registeredUser.Is_verified,
	})

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
