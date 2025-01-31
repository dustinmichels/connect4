// Usage:
// 	go run server/main.go
// 	go run server/main.go --port 50052

package server

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
	msg := fmt.Sprintf("Hello, %s! Welcome to Connect4!", req.PlayerName)
	return &pb.JoinResponse{
		PlayerId: "player1",
		Message:  msg,
	}, nil
}

func Run(ngrok bool) {

	// Parse flag and start the server
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Optionally start ngrok
	if ngrok {
		launchNgrok()
	}

	// Start the gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterConnect4GameServer(grpcServer, &server{})
	log.Printf("server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func launchNgrok() {
	fmt.Println("Starting ngrok tunnel...")
	tunnelURL, err := startNgrokTunnel(fmt.Sprintf("%d", *port))
	if err != nil {
		log.Fatalf("Error starting ngrok: %v", err)
	}
	fmt.Printf("gRPC Server is publicly accessible at: %s\n", tunnelURL)
}
