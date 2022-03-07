package main

import (
	"log"
	"simple-customer-service/config"
	"simple-customer-service/connection"
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

	connection.OpenDatabase()
}
