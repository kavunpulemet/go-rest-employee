package utils

import (
	"encoding/json"
	"net/http"
)

type StatusResponse struct {
	Status string `json:"status"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(ctx MyContext, w http.ResponseWriter, message string, statusCode int) {
	ctx.Logger.Error(message)

	errRes := ErrorResponse{Message: message}

	jsonErrRes, err := json.Marshal(errRes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonErrRes)
}
