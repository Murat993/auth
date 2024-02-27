package auth

import (
	"github.com/Murat993/auth/internal/client/db"
	"github.com/Murat993/auth/internal/repository"
	def "github.com/Murat993/auth/internal/service"
)

var _ def.AuthService = (*serverAuth)(nil)

type serverAuth struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

func NewService(userRepository repository.UserRepository, txManager db.TxManager) *serverAuth {
	return &serverAuth{
		userRepository: userRepository,
		txManager:      txManager,
	}
}
