package main

import (
	pb "grpc-modes/bidirectional-streaming/proto"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

var robotMsgMap = map[string]string{
	"hello":             "hi, i am a robot, welcom~",
	"how are you":       "fine, thank you, and you?",
	"i am fine too":     "boom~",
	"what is your name": "my name is robot gopher",
	"":                  "you need say something...",
}

type bidirectionalStreamServer struct{}

func (s *bidirectionalStreamServer) GetChatMsg(stream pb.ChatService_GetChatMsgServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		msg, has := robotMsgMap[req.Msg]

		if !has {
			msg = "i don't know what u are saying..."
		}
		if req.Msg == "bye" {
			break
		}
		var chatMsg pb.ChatMessage = pb.ChatMessage{
			Msg: msg,
		}
		err = stream.Send(&chatMsg)
		if err != nil {
			return err
		}
		log.Printf("[RECEVIED REQUEST]: %v\n", req)
	}
	return nil
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
	pb.RegisterChatServiceServer(grpcServer, &bidirectionalStreamServer{})
	grpcServer.Serve(listener)
}
