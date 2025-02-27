package main

import (
	"context"
	protoc "grpc/protoc"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client protoc.ExampleClient

func main() {
	//connection to internal gpc server
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client = protoc.NewExampleClient(conn)

	r := gin.Default()
	r.GET("/sent-message/:message", clientConnection)
	r.Run(":8080")
}

func clientConnection(c *gin.Context) {
	message := c.Param("message")

	req := &protoc.HelloRequest{SomeString: message}

	client.ServeReply(context.TODO(), req)

	c.JSON(http.StatusOK, gin.H{
		"message": "message sent successfully to serve:= " + message,
	})
}
