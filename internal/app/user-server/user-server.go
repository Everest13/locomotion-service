package user_server

import (
	"google.golang.org/grpc"
	"locomotion-service/internal/service/user"
	desc "locomotion-service/pkg/user-service"
)

type Implementation struct {
	desc.UnimplementedUserServiceServer
	userService *user.Service
}

func NewServer(userService *user.Service) *Implementation {
	return &Implementation{
		userService: userService,
	}
}

func (i *Implementation) RegisterServer(s *grpc.Server) {
	desc.RegisterUserServiceServer(s, i)
}
