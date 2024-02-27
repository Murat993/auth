package auth

import (
	"context"
	desc "github.com/Murat993/auth/pkg/auth_v1"
	"log"
)

func (i Implementation) GetRefreshToken(ctx context.Context, req *desc.GetRefreshTokenRequest) (*desc.GetRefreshTokenResponse, error) {
	refreshToken, err := i.authService.GetRefreshToken(ctx, req)

	if err != nil {
		return nil, err
	}

	log.Printf("GetRefreshToken")

	return &desc.GetRefreshTokenResponse{RefreshToken: refreshToken}, nil
}
