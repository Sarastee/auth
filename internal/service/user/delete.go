package user

import (
	"context"
	"fmt"
)

// Delete ...
func (s *Service) Delete(ctx context.Context, userID int64) error {
	if err := s.userRepo.Delete(ctx, userID); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
