package user_server

import (
	"context"
	desc "locomotion-service/pkg/user-service"
)

func (i *Implementation) GetUser(ctx context.Context, req *desc.GetUserRequest) (*desc.User, error) {
	user, err := i.userService.GetUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return convertUserToPB(user), nil
}
