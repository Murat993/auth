package converter

import (
	"github.com/Murat993/auth/internal/dto"
	modelRepo "github.com/Murat993/auth/internal/repository/user/entity"
)

func ToUserFromRepo(user *modelRepo.User) *dto.User {
	return &dto.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
