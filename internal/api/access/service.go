package access

import (
	"github.com/Murat993/auth/internal/service"
	desc "github.com/Murat993/auth/pkg/access_v1"
)

type Implementation struct {
	desc.UnimplementedAccessV1Server
	accessService service.AccessService
}

func NewImplementation(accessService service.AccessService) *Implementation {
	return &Implementation{
		accessService: accessService,
	}
}
