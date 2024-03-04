package user

import (
	"github.com/sarastee/auth/internal/repository"
	"github.com/sarastee/auth/internal/service"
)

var _ service.UserService = (*Service)(nil)

// Service ...
type Service struct {
	userRepo repository.UserRepository
}

// NewService ...
func NewService(userRepository repository.UserRepository) *Service {
	return &Service{
		userRepo: userRepository,
	}
}
