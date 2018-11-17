package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:8888"
)

// 定义helloService并实现约定的接口
type helloService struct{}

var HelloService = helloService{}

/*
 rpc SayHello (HelloRequest) returns (HelloReply) {}
*/
func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := new(pb.HelloReply)
	resp.Message = "grpc test:Hello " + in.Name + "."
	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloService
	pb.RegisterGreeterServer(s, HelloService)

	grpclog.Println("Listen on " + Address)

	s.Serve(listen)

}
