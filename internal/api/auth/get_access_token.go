package auth

import (
	"context"
	desc "github.com/Murat993/auth/pkg/auth_v1"
	"log"
)

func (i Implementation) GetAccessToken(ctx context.Context, req *desc.GetAccessTokenRequest) (*desc.GetAccessTokenResponse, error) {
	accessToken, err := i.authService.GetAccessToken(ctx, req)

	if err != nil {
		return nil, err
	}

	log.Printf("GetAccessToken")

	return &desc.GetAccessTokenResponse{AccessToken: accessToken}, nil
}
