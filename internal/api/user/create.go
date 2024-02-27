package user

import (
	"context"
	"github.com/Murat993/auth/internal/converter"
	desc "github.com/Murat993/auth/pkg/user_v1"
	"log"
)

func (i Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.userService.Create(ctx, converter.ToUserFromDescCreate(req.GetUseCreate()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted user with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
