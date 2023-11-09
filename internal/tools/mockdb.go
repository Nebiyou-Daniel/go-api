package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]*LoginDetails{
	"nebiyou": {
		AuthToken: "1234567890",
		Username: "nebiyou",
	},
	"john": {
		AuthToken: "0987654321",
		Username: "john",
	},
}

var mockCoinDetails = map[string]*CoinDetails{
	"nebiyou": {
		Coins: 100,
		Username: "nebiyou",
	},
	"john": {
		Coins: 200,
		Username: "john",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails{
	time.Sleep(time.Second * 1)

	var clientData *LoginDetails
	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}
	return clientData

}

func (d *mockDB) GetUserCoins(username string) *CoinDetails{
	time.Sleep(time.Second * 1)

	var clientData *CoinDetails
	clientData, ok := mockCoinDetails[username]

	if !ok {
		return nil
	}
	return clientData
}

func (d *mockDB) SetupDatabase() error{
	return nil
}
