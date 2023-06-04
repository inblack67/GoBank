package main

import "gorm.io/gorm"

type Storage interface {
	createAccount(*Account) (*Account, error)
	getAccount(id int) (*Account, error)
}

type PostgresStore struct {
	db *gorm.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) createAccount(acc *Account) (*Account, error) {
	res := s.db.Create(acc)
	if res.Error != nil {
		return nil, res.Error
	}
	return acc, nil
}

func (s *PostgresStore) getAccount(id int) (*Account, error) {
	acc := &Account{}
	res := s.db.Find(acc, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return acc, nil
}
