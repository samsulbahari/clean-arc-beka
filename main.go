package main

import (
	"fmt"
	lumber "grpc-beka/infa/log"
	"grpc-beka/infa/postgresql"
	"net"

	"google.golang.org/grpc"

	"grpc-beka/infa/grpc/cleanservice"
)

func main() {
	log := lumber.New()

	db_handler, err := postgresql.NewPool()
	if err != nil {
		log.WriteLn(err)
		return
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.WriteLn(err)
		return
	}

	grpc_service := cleanservice.NewGRPCService(log, db_handler)

	grpc := grpc.NewServer()
	cleanservice.RegisterRefSysAgamaServerServer(grpc, grpc_service)

	grpc.Serve(lis)
}
