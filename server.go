package main

import (
	"context"
	"echo/pb"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Echo(ctx context.Context, req *pb.EchoRequest) (res *pb.EchoResponse, err error) {
	fmt.Printf("Echo: %s\n", req.Echo)
	res = new(pb.EchoResponse)
	res.Echo = req.Echo
	return res, err
}

func main() {
	port := "5555"
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	fmt.Printf("Starting EchoServer on port %s\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
