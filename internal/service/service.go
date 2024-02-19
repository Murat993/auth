package service

import (
	"context"
	"github.com/Murat993/auth/internal/dto"
)

type UserService interface {
	Create(ctx context.Context, user *dto.UserCreate) (string, error)
	Get(ctx context.Context, id string) (*dto.User, error)
	Update(ctx context.Context, user *dto.UserUpdate) (string, error)
	Delete(ctx context.Context, id string) error
}
