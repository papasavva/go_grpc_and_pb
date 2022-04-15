package main

import (
	"context"
	"fmt"

	pb "github.com/papasavva/go_grpc_and_pb/gen/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	client := pb.NewTestApiClient(conn)

	resp, err := client.Echo(context.Background(), &pb.ResponseRequest{Msg: "Hello there!"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
	fmt.Println(resp.Msg)

	createdUser, err := client.GetUser(context.Background(), &pb.UserRequest{
		Id: "32",
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(createdUser)
}
