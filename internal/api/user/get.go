package user

import (
	"context"
	"errors"

	"github.com/sarastee/auth/internal/converter"
	"github.com/sarastee/auth/internal/repository"
	"github.com/sarastee/auth/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Get user by params
func (i *Implementation) Get(ctx context.Context, request *user_v1.GetRequest) (*user_v1.GetResponse, error) {
	user, err := i.userService.Get(ctx, request.Id)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrUserNotFound):
			return nil, status.Error(codes.NotFound, repository.ErrUserNotFound.Error())
		default:
			return nil, status.Error(codes.Internal, "User error")
		}
	}

	return converter.ToGetResponseFromServiceUser(user), nil
}
