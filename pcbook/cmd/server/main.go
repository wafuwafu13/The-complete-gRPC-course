package main

import (
	"fmt"
	"flag"
	"log"
	"net"
	"github.com/wafuwafu13/The-complete-gRPC-course/pb"
	"github.com/wafuwafu13/The-complete-gRPC-course/service"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)

	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("can not start server:", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("can not start server:", err)
	}
}
