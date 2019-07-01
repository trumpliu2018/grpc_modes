package main

import (
	"google.golang.org/grpc"
	"log"
	pb "grpc-mode/simple/proto"
	"context"
)

func main() {
	addr := "0.0.0.0:2333"

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect server error: %v", err)
	}
	defer conn.Close()

	grpcClient := pb.NewUserServiceClient(conn)

	req := pb.UserRequest{ID: 2}
	resp, err := grpcClient.GetUserInfo(context.Background(), &req)
	if err != nil {
		log.Fatalf("recevie resp error: %v", err)
	}

	log.Printf("[RECEIVED RESPONSE]: %v\n", resp)
}
