package main

import "log"

func main() {
	storage, err := NewPostgresStore()
	if err != nil {
		log.Fatal("Error connecting to db[postgres]=", err)
	}
	server := NewAPIServer(":3000", storage)
	server.Run()
}
