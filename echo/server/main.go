package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/yumuranaoki/grpc-practice/echo/proto"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("[echo] ")
}

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen")
	}

	srv := grpc.NewServer()
	pb.RegisterEchoServiceServer(srv,
		&echoService{})

	log.Printf("start server on port%s\n", port)
	if err = srv.Serve(lis); err != nil {
		log.Printf("failed to serve")
	}
}
