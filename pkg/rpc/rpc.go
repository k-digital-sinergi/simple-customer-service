package rpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"simple-customer-service/config"
	customer "simple-customer-service/pkg/rpc/customer/proto"
)

func Start() {
	server := grpc.NewServer()

	reflection.Register(server)

	customer.RegisterCustomerServiceServer(server, customer.NewRpcServer())

	addr := config.Env.CustomerGrpcAddr
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("GRPC server started on [::]:%s", addr))

	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
