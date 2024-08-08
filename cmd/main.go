package main

import (
	"fmt"
	"home-manager/server/internal/application"
	"home-manager/server/internal/infrastructure/db"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db := db.Init()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", "50051"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	application.SetupServers(s, &db)
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
