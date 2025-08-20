package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	_ "user-srv/cmd/inits"
	__ "user-srv/cmd/proto"
	"user-srv/handler"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	__.RegisterUserServer(s, &handler.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
