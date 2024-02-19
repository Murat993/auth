package user

import (
	"context"
	"github.com/Murat993/auth/internal/converter"
	desc "github.com/Murat993/auth/pkg/user_v1"
	"log"
)

func (i Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	userObj, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %s, role: %s, name: %s, email: %s, created_at: %v, updated_at: %v\n", userObj.ID, userObj.Role, userObj.Name, userObj.Role, userObj.CreatedAt, userObj.UpdatedAt)

	return &desc.GetResponse{
		User: converter.ToDescUserFromDto(userObj),
	}, nil
}
