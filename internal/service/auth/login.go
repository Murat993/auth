package auth

import (
	"context"
	"github.com/Murat993/auth/internal/config/env"
	"github.com/Murat993/auth/internal/dto"
	"github.com/Murat993/auth/internal/utils"
	"github.com/pkg/errors"
	"os"
)

func (s serverAuth) Login(ctx context.Context, login *dto.UserLogin) (string, error) {
	// Лезем в базу или кэш за данными пользователя
	// Сверяем хэши пароля
	var user *dto.User
	var errTx error
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		user, errTx = s.userRepository.GetByUsername(ctx, login)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	refreshToken, err := utils.GenerateToken(user,
		[]byte(os.Getenv(env.RefreshTokenSecretKey)),
		env.RefreshTokenExpiration,
	)

	if errTx != nil {
		return "", errors.New("failed to generate token")
	}

	return refreshToken, nil
}
