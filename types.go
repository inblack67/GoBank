package main

import "math/rand"

type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Number   int64  `json:"number"`
	Balance  int64  `json:"balance"`
}

func NewAccount(username string) *Account {
	return &Account{
		ID:       rand.Intn(100000),
		Username: username,
		Number:   int64(rand.Intn(1000000)),
		// Balance - by default 0
	}
}
