package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/papasavva/go_grpc_and_pb/gen/proto"
	"google.golang.org/grpc"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (t *testApiServer) Echo(ctx context.Context, in *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	return in, nil
}

func (t *testApiServer) GetUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		Id:         "32",
		Name:       "Boby",
		Age:        0,
		Email:      "",
		Registered: false,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error tcp")
		return
	}
	grpcServer := grpc.NewServer()

	pb.RegisterTestApiServer(grpcServer, &testApiServer{})

	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Println("Error serve")

		return
	}
}
