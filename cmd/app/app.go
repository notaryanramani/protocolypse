package app

import (
	"google.golang.org/grpc"
)

type Application struct {
	Server *grpc.Server
}

func NewApplication() *Application {
	return &Application{
		Server: grpc.NewServer(),
	}
}