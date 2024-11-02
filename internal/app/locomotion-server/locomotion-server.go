package locomotion_server

import (
	"google.golang.org/grpc"
	"locomotion-service/internal/service/locomotion"
	desc "locomotion-service/pkg/locomotion-service"
)

type Implementation struct {
	desc.UnimplementedLocomotionServiceServer
	locomotionService *locomotion.Service
}

func NewServer(locomotionService *locomotion.Service) *Implementation {
	return &Implementation{locomotionService: locomotionService}
}

func (i *Implementation) RegisterServer(s *grpc.Server) {
	desc.RegisterLocomotionServiceServer(s, i)
}
