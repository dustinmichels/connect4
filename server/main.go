// Usage:
// 	go run server/main.go
// 	go run server/main.go --port 50052

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "connect4/proto" // Importing the generated Protobuf code with alias "pb"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedConnect4GameServer
}

func (s *server) JoinGame(ctx context.Context, req *pb.JoinRequest) (*pb.JoinResponse, error) {
	return &pb.JoinResponse{
		PlayerId: "player1",
		Message:  "Welcome to Connect4!",
	}, nil
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterConnect4GameServer(grpcServer, &server{})

	log.Printf("server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
