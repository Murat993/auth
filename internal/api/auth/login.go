package auth

import (
	"context"
	"github.com/Murat993/auth/internal/dto"
	descAuth "github.com/Murat993/auth/pkg/auth_v1"
	"github.com/Murat993/platform_common/pkg/sys/validate"
	"log"
)

func (i Implementation) Login(ctx context.Context, req *descAuth.LoginRequest) (*descAuth.LoginResponse, error) {
	//err := validate.Validate(
	//	ctx,
	//	validateID(req.GetId()),
	//	otherValidateID(req.GetId()),
	//)

	refreshToken, err := i.authService.Login(ctx, &dto.UserLogin{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	})

	if err != nil {
		return nil, err
	}

	log.Printf("login username")

	return &descAuth.LoginResponse{RefreshToken: refreshToken}, nil
}

func validateID(id int64) validate.Condition {
	return func(ctx context.Context) error {
		if id <= 0 {
			return validate.NewValidationErrors("id must be greater than 0")
		}

		return nil
	}
}

func otherValidateID(id int64) validate.Condition {
	return func(ctx context.Context) error {
		if id <= 100 {
			return validate.NewValidationErrors("id must be greater than 100")
		}

		return nil
	}
}
