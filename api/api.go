package api

import (
	"encoding/json"
	"net/http"
)

// Coin balance Params
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	StatusCode int
	Balance int64 
}

// Error
type Error struct {
	StatusCode int
	Message string
}

func writeError(w http.ResponseWriter, message string, statusCode int){
	resp := Error{
		StatusCode: statusCode,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(resp)
}

var(
	RequestErrorHandler = func(w http.ResponseWriter, err error){
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter){
		writeError(w, "An Unexpected Error Occurred.", http.StatusInternalServerError)
	}
)
