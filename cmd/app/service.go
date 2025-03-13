package app

import (
	"context"
	"log"

	pb "github.com/notaryanramani/protocolypse/pb/api/proto"
)

type ChatService struct {
	pb.UnimplementedChatServiceServer
}

func NewChatService() *ChatService {
	return &ChatService{}
}

func (s *ChatService) SendMessage(ctx context.Context, in *pb.MessageRequest) (*pb.MessageResponse, error) {
	log.Printf("Received message from %s: %s", in.User, in.Message)
	return &pb.MessageResponse{Message: in.Message}, nil
}
