package main

import (
	"grpc/proto"
	"net"

	"google.golang.org/grpc"
)

func main() {
	addr := "localhost:50051"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	serve := grpc.NewServer()
	proto.RegisterPictureConServiceServer(serve, &proto.UnimplementedPictureConServiceServer{})
	err = serve.Serve(listener)
	if err != nil {
		panic(err)
	}
}
