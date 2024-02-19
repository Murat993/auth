package user

import (
	"context"
	"github.com/Murat993/auth/internal/converter"
	desc "github.com/Murat993/auth/pkg/user_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
)

func (i Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {
	id, err := i.userService.Update(ctx, converter.ToUserFromDescUpdate(req.GetId(), req.GetUserUpdate()))
	if err != nil {
		return nil, err
	}

	log.Printf("updated user with id: %s", id)

	return &empty.Empty{}, nil
}
