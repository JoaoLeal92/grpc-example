package main

import (
	"context"
	"log"
	"net"

	pb "github.com/JoaoLeal92/grcp-example/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedExampleGrpcServer
}

func (s *server) SayHello(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	return &pb.Response{
		Message: "You have successfully implemented a gRCP connection!",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("error connecting to port")
	}

	s := grpc.NewServer()
	pb.RegisterExampleGrpcServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatal("error serving: ", err.Error())
	}

}
