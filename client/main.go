package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "connect4/proto" // Import the generated Protobuf package

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "38.42.221.206:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.

	fmt.Printf("Connecting to %s\n", *addr)

	// conn, err := grpc.Dial("192.168.1.10:50051", grpc.WithInsecure()) // Replace with your server's IP
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewConnect4GameClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.JoinGame(ctx, &pb.JoinRequest{PlayerName: "Player"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s", r.GetMessage())

}
