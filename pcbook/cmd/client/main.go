package main

import (
	//"fmt"
	"time"
	"context"
	"flag"
	"log"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"github.com/wafuwafu13/The-complete-gRPC-course/pb"
	"github.com/wafuwafu13/The-complete-gRPC-course/sample"

)

func main() {
	serverAdress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAdress)
	
	conn, err := grpc.Dial(*serverAdress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()
	// laptop.Id = "31d5c727-9c21-4af8-9c48-b4f6adca6810"
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()


	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Printf("laptop already exists")
		} else {
			log.Fatal("cannot create laptop: ", err)
		}
		return
	}

	log.Printf("crated laptop with id: %s", res.Id)
}
