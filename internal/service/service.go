package service

import (
	"context"
	"github.com/Murat993/auth/internal/dto"
	descCheck "github.com/Murat993/auth/pkg/access_v1"
	descAuth "github.com/Murat993/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService interface {
	Create(ctx context.Context, user *dto.UserCreate) (int64, error)
	Get(ctx context.Context, id int64) (*dto.User, error)
	Update(ctx context.Context, user *dto.UserUpdate) (int64, error)
	Delete(ctx context.Context, id int64) error
}

type AuthService interface {
	Login(ctx context.Context, login *dto.UserLogin) (string, error)
	GetRefreshToken(ctx context.Context, req *descAuth.GetRefreshTokenRequest) (string, error)
	GetAccessToken(ctx context.Context, req *descAuth.GetAccessTokenRequest) (string, error)
}

type AccessService interface {
	Check(ctx context.Context, req *descCheck.CheckRequest) (*emptypb.Empty, error)
}
