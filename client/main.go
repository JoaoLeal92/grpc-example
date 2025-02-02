package main

import (
	"context"
	"log"
	"time"

	pb "github.com/JoaoLeal92/grcp-example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error creating client: ", err.Error())
	}
	defer conn.Close()

	c := pb.NewExampleGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.SayHello(ctx, &pb.Request{})
	if err != nil {
		log.Fatal("error getting response from server: ", err.Error())
	}

	log.Printf("got server response: %v", res.GetMessage())
}
