package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	// The error message
	// required: false
	Error string `json:"error"`

	// The error code
	// required: false
	ErrorCode int `json:"errorCode"`
}

func ErrorResponseFunc(w http.ResponseWriter, r *http.Request, err error) {
	//fmt.Printf("ERROR: %s\n", err.Error())

	errorResponse := ErrorResponse{Error: err.Error()}

	switch {
	case IsErrUnauthorized(err):
		errorResponse.ErrorCode = http.StatusUnauthorized

	default:

		errorResponse.Error = "internal server error"
		errorResponse.ErrorCode = http.StatusInternalServerError
	}
	//fmt.Printf("API ERROR: %s\n", err.Error())
	SetResponseHeader(w, "Content-Type", "application/json")
	data, err := json.Marshal(errorResponse)
	if err != nil {
		data = []byte("{}")
	}

	w.WriteHeader(errorResponse.ErrorCode)
	_, _ = w.Write(data)
}

func StringResponse(w http.ResponseWriter, message string) {
	SetResponseHeader(w, "Content-Type", "text/plain")
	_, _ = fmt.Fprint(w, message)
}

func JsonStringResponse(w http.ResponseWriter, code int, message string) { //nolint:unparam
	SetResponseHeader(w, "Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}

func JsonBytesResponse(w http.ResponseWriter, code int, json []byte) { //nolint:unparam
	SetResponseHeader(w, "Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(json)
}

func SetResponseHeader(w http.ResponseWriter, key string, value string) { //nolint:unparam
	header := w.Header()
	if header == nil {
		return
	}
	header.Set(key, value)
}
