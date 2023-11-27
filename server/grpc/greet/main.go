package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/harrylee2015/harry_tools/server/grpc/greet/types"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime" // 注意v2版本
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	types.UnimplementedGreeterServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) Greeter(ctx context.Context, in *types.GreeterRequest) (*types.GreeterResponse, error) {
	return &types.GreeterResponse{Message: in.GetMsg() + " world"}, nil
}

func (s *server) GetGreet(ctx context.Context, in *types.GreeterRequest) (*types.GreeterResponse, error) {
	return &types.GreeterResponse{Message: in.GetMsg() + " world"}, nil
}
func main() {

	ctx := context.Background()
	// Create a listener on TCP port
	listen, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// 创建一个gRPC server对象
	s := grpc.NewServer()
	// 注册Greeter service到server
	types.RegisterGreeterServer(s, &server{})
	// 8080端口启动gRPC Server
	log.Println("Serving gRPC on :8080")
	go func() {
		err := s.Serve(listen)

		if err != nil {
			log.Fatalln(err)
		}

	}()
	fmt.Println("start api gateway...")
	time.Sleep(5 * time.Second)

	// 创建一个连接到我们刚刚启动的 gRPC 服务器的客户端连接
	// gRPC-Gateway 就是通过它来代理请求（将HTTP请求转为RPC请求）
	conn, err := grpc.DialContext(
		ctx,
		"localhost:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	fmt.Println("start api gateway...1")
	gwmux := runtime.NewServeMux()
	// 注册Greeter
	err = types.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	fmt.Println("start api gateway...2")

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}
	// 8090端口提供gRPC-Gateway服务
	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
