package access

import (
	"github.com/Murat993/auth/internal/client/db"
	"github.com/Murat993/auth/internal/repository"
	def "github.com/Murat993/auth/internal/service"
)

var _ def.AccessService = (*serverAccess)(nil)

type serverAccess struct {
	accessRepository repository.AccessRepository
	txManager        db.TxManager
}

func NewService(accessRepository repository.AccessRepository, txManager db.TxManager) *serverAccess {
	return &serverAccess{
		accessRepository: accessRepository,
		txManager:        txManager,
	}
}
