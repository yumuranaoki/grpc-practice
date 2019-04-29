package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/yumuranaoki/grpc-practice/downloader/proto"
)

const port = ":50051"

func main() {
	log.SetFlags(0)
	log.SetPrefix("[downloader] ")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on %v\n", port)
	}

	srv := grpc.NewServer()
	pb.RegisterFileServiceServer(srv, &fileService{})
	log.Printf("start server on ports %s\n", port)
	if err := srv.Serve(lis); err != nil {
		log.Printf("failed to serve %v\n", err)
	}
}
