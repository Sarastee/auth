package user

import (
	"context"

	"github.com/sarastee/auth/internal/model"
)

func (s *Service) Delete(ctx context.Context, userID model.UserID) error {
	if err := s.userRepo.Delete(ctx, userID); err != nil {
		return err
	}

	return nil
}
