package main

import (
	"fmt"
	"log"
	"net"

	"github.com/spinel/gophkeeper-storage-svc/pkg/config"
	"github.com/spinel/gophkeeper-storage-svc/pkg/db"
	"github.com/spinel/gophkeeper-storage-svc/pkg/pb"
	"github.com/spinel/gophkeeper-storage-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	fmt.Println("Storage Svc on", c.Port)

	s := services.Server{
		H: h,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterStorageServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
