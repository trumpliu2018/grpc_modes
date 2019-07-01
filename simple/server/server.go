package main

import (
	"context"
	pb "grpc-modes/simple/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var users = map[int32]pb.UserResponse{
	1: {Name: "Dennis MacAlistair Ritchie", Age: 70},
	2: {Name: "Ken Thompson", Age: 75},
	3: {Name: "Rob Pike", Age: 62},
}

type simpleServer struct{}

func (s *simpleServer) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	if user, ok := users[req.ID]; ok {
		resp = &user
	}
	log.Printf("[RECEVIED REQUEST]: %v\n", req)
	return
}

func main() {
	addr := "0.0.0.0:2333"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	} else {
		log.Println("server listen: ", addr)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &simpleServer{})
	grpcServer.Serve(listener)
}
