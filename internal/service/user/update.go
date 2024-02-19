package user

import (
	"context"
	"github.com/Murat993/auth/internal/dto"
)

func (s serv) Update(ctx context.Context, userUpdate *dto.UserUpdate) (string, error) {
	var id string
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.userRepository.Update(ctx, userUpdate)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.userRepository.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return id, nil
}
