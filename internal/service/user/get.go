package user

import (
	"context"
	"github.com/Murat993/auth/internal/dto"
)

func (s serv) Get(ctx context.Context, id string) (*dto.User, error) {
	user, err := s.userRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
