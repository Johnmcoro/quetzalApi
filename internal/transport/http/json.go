package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

const (
	INTERNAL_SERVER_ERROR = "Internal Server Error"
)

//TODO - move to error package
type ApiError struct {
	Status        int         `json:"status"`
	Timestamp     *time.Time  `json:"timestamp,omitempty"`
	Message       string      `json:"message"`
	Errors        interface{} `json:"errors,omitempty"`
	InternalError string      `json:"-"`
}

type errorResponse struct {
	Error ApiError `json:"error"`
}

func (a *ApiError) Error() string {
	json, err := json.Marshal(errorResponse{Error: *a})
	if err != nil {
		return err.Error()
	}
	return string(json)
}

type apiResponse struct {
	Message interface{} `json:"message"`
}

func WriteJSONResponse(w http.ResponseWriter, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json, err := json.Marshal(apiResponse{Message: payload})
	if err != nil {
		w.Write([]byte("Response Error"))
		return
	}
	w.Write(json)
}

func WriteJSONError(w http.ResponseWriter, err error) {
	var apiError *ApiError
	ok := errors.As(err, &apiError)
	if !ok {
		print("err\n", err.Error())
		apiError.Message = INTERNAL_SERVER_ERROR
		apiError.Status = http.StatusInternalServerError
		apiError.InternalError = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte(apiError.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(apiError.Status)
	w.Write([]byte(apiError.Error()))
}
