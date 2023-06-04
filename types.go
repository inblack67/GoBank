package main

import "math/rand"

type Account struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Number   int64  `json:"number"`
	Balance  int64  `json:"balance"`
}

type CreateAccountReq struct {
	Username string `json:"username"`
}

func NewAccount(username string) *Account {
	return &Account{
		Username: username,
		Number:   int64(rand.Intn(1000000)),
		// Balance - by default 0
	}
}
