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

// Create user by params
func (i *Implementation) Create(ctx context.Context, request *user_v1.CreateRequest) (*user_v1.CreateResponse, error) {
	userCreate := converter.ToServiceUserCreateFromCreateRequest(request)
	userID, err := i.userService.Create(ctx, userCreate)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrEmailIsTaken):
			return nil, status.Error(codes.AlreadyExists, repository.ErrEmailIsTaken.Error())
		default:
			return nil, status.Error(codes.Internal, "Error")
		}
	}

	return &user_v1.CreateResponse{Id: userID}, nil
}
