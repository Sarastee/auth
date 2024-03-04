package user

import (
	"context"
	"fmt"

	"github.com/sarastee/auth/internal/model"
)

// Get ...
func (s *Service) Get(ctx context.Context, userID int64) (*model.User, error) {
	user, err := s.userRepo.Get(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}
