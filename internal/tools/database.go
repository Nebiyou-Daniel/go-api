package tools

import (
	log "github.com/sirupsen/logrus"
)

// Database Collections

// Login Details
type LoginDetails struct {
	AuthToken string
	Username string
}

// Coin Balance
type CoinDetails struct {
	Username string
	Coins int64
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error){

	var database DatabaseInterface = &mockDB{}
	var err error = database.SetupDatabase()
	 
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}
