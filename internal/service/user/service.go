package user

import (
	"github.com/sarastee/auth/internal/repository"
	"github.com/sarastee/auth/internal/service"
)

var _ service.UserService = (*Service)(nil)

type Service struct {
	userRepo repository.UserRepository
}

func NewService(userRepository repository.UserRepository) *Service {
	return &Service{
		userRepo: userRepository,
	}
}
