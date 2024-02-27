package converter

import (
	"github.com/Murat993/auth/internal/dto"
	desc "github.com/Murat993/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserFromDescCreate(userCreate *desc.UserCreate) *dto.UserCreate {
	return &dto.UserCreate{
		Name:            userCreate.GetName(),
		Email:           userCreate.GetPassword(),
		Password:        userCreate.GetPassword(),
		PasswordConfirm: userCreate.GetPasswordConfirm(),
		Role:            userCreate.GetRole(),
	}
}

func ToDescUserFromDto(userDto *dto.User) *desc.User {
	var updatedAt *timestamppb.Timestamp
	if userDto.UpdatedAt.Valid {
		updatedAt = timestamppb.New(userDto.UpdatedAt.Time)
	}

	return &desc.User{
		Id:        userDto.ID,
		Name:      userDto.Name,
		Email:     userDto.Email,
		Role:      userDto.Role,
		CreatedAt: timestamppb.New(userDto.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToUserFromDescUpdate(ID int64, userUpdate *desc.UserUpdate) *dto.UserUpdate {
	return &dto.UserUpdate{
		ID:    ID,
		Name:  userUpdate.GetName(),
		Email: userUpdate.GetEmail(),
	}
}
