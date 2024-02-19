package repository

import (
	"context"
	"github.com/Murat993/auth/internal/dto"
)

type UserRepository interface {
	Create(ctx context.Context, userCreate *dto.UserCreate) (string, error)
	Get(ctx context.Context, id string) (*dto.User, error)
	Update(ctx context.Context, userUpdate *dto.UserUpdate) (string, error)
	Delete(ctx context.Context, id string) error
}
