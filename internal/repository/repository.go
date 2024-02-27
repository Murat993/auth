package repository

import (
	"context"
	"github.com/Murat993/auth/internal/dto"
	"github.com/Murat993/auth/pkg/user_v1"
)

type UserRepository interface {
	Create(ctx context.Context, userCreate *dto.UserCreate) (int64, error)
	Get(ctx context.Context, id int64) (*dto.User, error)
	Update(ctx context.Context, userUpdate *dto.UserUpdate) (int64, error)
	Delete(ctx context.Context, id int64) error
	GetByUsername(ctx context.Context, login *dto.UserLogin) (*dto.User, error)
}

type AccessRepository interface {
	Check(ctx context.Context) (map[string]user_v1.Role, error)
}
