package api

import (
	// "encoding/json"
	// "net/http"
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
