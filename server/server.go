package main

import (
	"context"
	"errors"
	"fmt"
	proto "grpc/protoc"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedExampleServer
}

func main() {

	listener, tcpErr := net.Listen("tcp", ":9000")
	if tcpErr != nil {
		panic(tcpErr)
	}
	srv := grpc.NewServer()
	proto.RegisterExampleServer(srv, &server{})
	reflection.Register(srv)

	err := srv.Serve(listener)
	if err != nil {
		panic(err)
	}
}

func (s *server) ServeReply(c context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	fmt.Println("recieve request from client", req.SomeString)
	fmt.Println("hello from server")
	return &proto.HelloResponse{}, errors.New("")
}
