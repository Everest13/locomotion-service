package user_server

import (
	"context"
	desc "locomotion-service/pkg/user-service"
)

func (i *Implementation) CreateUser(ctx context.Context, req *desc.UserCreate) (*desc.CreateUserResponse, error) {
	user := convertUserFromPB(req)
	err := i.userService.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &desc.CreateUserResponse{Result: true}, nil
}
