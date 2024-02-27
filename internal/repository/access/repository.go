package user

import (
	"context"
	"github.com/Murat993/auth/internal/client/db"
	def "github.com/Murat993/auth/internal/repository"
	user_v1 "github.com/Murat993/auth/pkg/user_v1"
)

var _ def.AccessRepository = (*repo)(nil)

type repo struct {
	db db.Client
}

const (
	ExamplePath = "/auth_v1.AuthV1/Get"
)

func NewAccessRepository(db db.Client) def.AccessRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Check(_ context.Context) (map[string]user_v1.Role, error) {
	accessbileRoles := make(map[string]user_v1.Role)
	accessbileRoles[ExamplePath] = user_v1.Role_USER

	return accessbileRoles, nil
}
