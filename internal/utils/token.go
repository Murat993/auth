package utils

import (
	"github.com/Murat993/auth/internal/dto"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

func GenerateToken(user *dto.User, secretKey []byte, duration time.Duration) (string, error) {
	claims := dto.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
		Username: user.Name,
		Role:     user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func VerifyToken(tokenStr string, secretKey []byte) (*dto.UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&dto.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.Errorf("unexpected token signing method")
			}

			return secretKey, nil
		},
	)
	if err != nil {
		return nil, errors.Errorf("invalid token: %s", err.Error())
	}

	claims, ok := token.Claims.(*dto.UserClaims)
	if !ok {
		return nil, errors.Errorf("invalid token claims")
	}

	return claims, nil
}
