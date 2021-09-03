package rpc

import (
	"context"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type greetingServiceServer struct {
	UnimplementedGreetingServiceServer
}

func (s *greetingServiceServer) Greeting(ctx context.Context, request *HelloRequest) (*HelloResponse, error) {
	// Greeting方法
	name := request.GetName()
	response := &HelloResponse{Greeting: fmt.Sprintf("hello %v", name)}
	log.Info(fmt.Sprintf("%s has visted ...", name))
	return response, nil
}

func newServer() GreetingServiceServer {
	return &greetingServiceServer{}
}

func StartServer() {
	port := 10081
	flag.Parse()
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Error(err)
	}
	log.Info(fmt.Sprintf("grpc listening %d...", port))
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	RegisterGreetingServiceServer(grpcServer, newServer())
	err2 := grpcServer.Serve(listen)
	if err2 != nil {
		log.Error(err2)
		return
	}
}
