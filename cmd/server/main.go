package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"pcbook/pb"
	"pcbook/service"
)

func main() {
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("start server on port:%v", *port)

	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)
	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listner, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
	err = grpcServer.Serve(listner)
	if err != nil {
		log.Fatal("cannot start server: ", err)

	}
}
