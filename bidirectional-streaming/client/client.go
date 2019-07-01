package main

import (
	"bufio"
	"context"
	"fmt"
	pb "grpc-modes/bidirectional-streaming/proto"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
)

func main() {
	// gRPC 服务器的地址
	addr := "0.0.0.0:2333"

	// 不使用认证建立连接
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect server error: %v", err)
	}
	defer conn.Close()

	// 创建 gRPC 客户端实例
	grpcClient := pb.NewChatServiceClient(conn)
	stream, err := grpcClient.GetChatMsg(context.Background())
	if err != nil {
		log.Fatalf("receive stream error: %v", err)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your msg: ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		stream.Send(&pb.ChatMessage{Msg: text})
		resp, err := stream.Recv()
		if err != nil {
			log.Fatalf("resp error: %v", err)
		}
		log.Printf("[RECEIVED RESPONSE]: %v\n", resp)
	}
	// 向服务端发送数据流，并处理响应流

	// var i int32
	// for i = 1; i < 4; i++ {
	// 	stream.Send(&pb.UserRequest{ID: i})
	// 	time.Sleep(1 * time.Second)
	// 	resp, err := stream.Recv()
	// 	if err != nil {
	// 		log.Fatalf("resp error: %v", err)
	// 	}
	// 	log.Printf("[RECEIVED RESPONSE]: %v\n", resp) // 输出响应
	// }
}
