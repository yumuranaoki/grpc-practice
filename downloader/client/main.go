package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "github.com/yumuranaoki/grpc-practice/downloader/proto"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("[file] ")
}

func main() {
	target := "localhost:50051"
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn note connect: %s\n", err)
	}
	defer conn.Close()

	c := pb.NewFileServiceClient(conn)
	name := os.Args[1]
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stream, err := c.Download(ctx, &pb.FileRequest{Name: name})
	if err != nil {
		log.Fatalf("could not downloaded: %s\n", err)
	}
	var blob []byte
	for {
		c, err := stream.Recv()
		if err == io.EOF {
			log.Printf("donw %d bytes\n", len(blob))
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		blob = append(blob, c.GetData()...)
	}
	ioutil.WriteFile(name, blob, 0644)
}
