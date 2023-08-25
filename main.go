package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "optimus.dev.br/controller/proto/optimus.dev.br/controller"
)

type server struct {
	pb.UnimplementedControllerServer
}

func (s *server) SetApplication(ctx context.Context, req *pb.Application) (*pb.Result, error) {
	return &pb.Result{Status: true, Message: "Message resut"}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Falied to listen port 50051 [%v]\n", err)
	}

	srv := grpc.NewServer()
	pb.RegisterControllerServer(srv, &server{})
	fmt.Println("Server listening on 50051")
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Error to server [%v]", err)
	}
	fmt.Println("Ok")
}
