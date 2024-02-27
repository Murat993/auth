package auth

import (
	"context"
	"github.com/Murat993/auth/internal/dto"
	descAuth "github.com/Murat993/auth/pkg/auth_v1"
	"log"
)

func (i Implementation) Login(ctx context.Context, req *descAuth.LoginRequest) (*descAuth.LoginResponse, error) {
	refreshToken, err := i.authService.Login(ctx, &dto.UserLogin{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	})

	if err != nil {
		return nil, err
	}

	log.Printf("login username")

	return &descAuth.LoginResponse{RefreshToken: refreshToken}, nil
}
