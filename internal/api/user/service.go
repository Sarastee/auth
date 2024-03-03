package user

import (
	"github.com/sarastee/auth/internal/service"
	"github.com/sarastee/auth/pkg/user_v1"
)

type Implementation struct {
	user_v1.UnimplementedUserV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
