package auth

import (
	"context"
	"github.com/Murat993/auth/internal/config/env"
	"github.com/Murat993/auth/internal/dto"
	"github.com/Murat993/auth/internal/utils"
	descAuth "github.com/Murat993/auth/pkg/auth_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
)

func (s *serverAuth) GetRefreshToken(_ context.Context, req *descAuth.GetRefreshTokenRequest) (string, error) {
	claims, err := utils.VerifyToken(req.GetRefreshToken(), []byte(os.Getenv(env.RefreshTokenSecretKey)))
	if err != nil {
		return "", status.Errorf(codes.Aborted, "invalid refresh token")
	}

	// Можем слазать в базу или в кэш за доп данными пользователя

	refreshToken, err := utils.GenerateToken(&dto.User{
		Name: claims.Username,
		Role: claims.Role,
	},
		[]byte(os.Getenv(env.RefreshTokenSecretKey)),
		env.RefreshTokenExpiration,
	)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
