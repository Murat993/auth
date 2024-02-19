package user

import (
	"context"
	desc "github.com/Murat993/auth/pkg/user_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
)

func (i Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	err := i.userService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("deleted user with id: %s", req.GetId())

	return &empty.Empty{}, nil
}
