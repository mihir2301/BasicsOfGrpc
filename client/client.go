package main

import (
	"context"
	protoc "grpc/protoc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//connection to internal gpc server
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := protoc.NewExampleClient(conn)

	req := &protoc.HelloRequest{SomeString: "Hello from client"}

	client.ServeReply(context.TODO(), req)
}
