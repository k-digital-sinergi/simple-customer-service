package main

import (
	"log"
	"simple-customer-service/config"
	"simple-customer-service/db"
	"simple-customer-service/pkg/rpc"
)

func main() {
	rpc.Start()
}

func init() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatal("error load env")
	}

	db.OpenDatabase()
}
