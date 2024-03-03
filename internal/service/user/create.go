package user

import (
	"context"
	"fmt"

	"github.com/sarastee/auth/internal/model"
)

func (s *Service) Create(ctx context.Context, userCreate *model.UserCreate) (model.UserID, error) {
	userID, err := s.userRepo.Create(ctx, userCreate)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return userID, nil
}
