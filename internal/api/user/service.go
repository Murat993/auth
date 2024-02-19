package user

import (
	"github.com/Murat993/auth/internal/service"
	desc "github.com/Murat993/auth/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
