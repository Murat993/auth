package user

import (
	"context"
)

func (s serv) Delete(ctx context.Context, id string) error {
	err := s.userRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
