package main

import (
	"log"
	"net"

	pb "github.com/notaryanramani/protocolypse/pb/api/proto"
	"github.com/notaryanramani/protocolypse/cmd/app"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	application := app.NewApplication()
	chatServer := app.NewChatService()
	pb.RegisterChatServiceServer(application.Server, chatServer)

	if err := application.Server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

