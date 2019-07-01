package main

import (
	pb "grpc-modes/server-side-streaming/proto"
	"net"
	"log"
	"google.golang.org/grpc"
	"time"
)

// 模拟的数据库查询结果
var users = map[int32]pb.UserResponse{
	1: {Name: "Dennis MacAlistair Ritchie", Age: 70},
	2: {Name: "Ken Thompson", Age: 75},
	3: {Name: "Rob Pike", Age: 62},
	4: {Name: "trump", Age: 18},
	5: {Name: "trump2", Age: 18},
	6: {Name: "trump3", Age: 18},
	7: {Name: "trump4", Age: 18},
}

type serverSideStreamServer struct{}

// serverSideStreamServer 实现了 user.pb.go 中的 UserServiceServer 接口
func (s *serverSideStreamServer) GetUserInfo(req *pb.UserRequest, stream pb.UserService_GetUserInfoServer) error {
	// 响应流数据
	log.Printf("[RECEIVED REQUEST]: %v\n", req)
	for _, user := range users {
		time.Sleep(time.Duration(10 * time.Second))
		log.Print("calculating.....", user.Name)
		stream.Send(&user)
	}
	return nil
}

func main() {
	// 指定微服务的服务端监听地址
	addr := "0.0.0.0:2333"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	} else {
		log.Println("server listen: ", addr)
	}

	// 创建 gRPC 服务器实例
	grpcServer := grpc.NewServer()

	// 向 gRPC 服务器注册服务
	pb.RegisterUserServiceServer(grpcServer, &serverSideStreamServer{})

	// 启动 gRPC 服务器
	// 阻塞等待客户端的调用
	grpcServer.Serve(listener)
}
