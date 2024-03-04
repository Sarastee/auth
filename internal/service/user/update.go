package user

import (
	"context"
	"fmt"

	"github.com/sarastee/auth/internal/model"
)

// Update ...
func (s *Service) Update(ctx context.Context, userUpdate *model.UserUpdate) error {
	if err := s.userRepo.Update(ctx, userUpdate); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
