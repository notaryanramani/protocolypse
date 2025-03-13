package main

import (
	"context"
	"log"

	pb "github.com/notaryanramani/protocolypse/pb/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)
	_, err = client.SendMessage(context.Background(), &pb.MessageRequest{User: "Aryan", Message: "Hello, World!"})

	if err != nil {
		log.Fatalf("failed to send message: %v", err)
	}
}
