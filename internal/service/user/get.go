package user

import (
	"context"

	"github.com/sarastee/auth/internal/model"
)

func (s *Service) Get(ctx context.Context, userID model.UserID) (*model.User, error) {
	user, err := s.userRepo.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
