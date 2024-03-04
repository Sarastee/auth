package user

import (
	"context"

	"github.com/sarastee/auth/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete user by params
func (i *Implementation) Delete(ctx context.Context, request *user_v1.DeleteRequest) (*emptypb.Empty, error) {
	err := i.userService.Delete(ctx, request.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "Delete error")
	}

	return &emptypb.Empty{}, nil
}
