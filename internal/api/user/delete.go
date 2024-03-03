package user

import (
	"context"

	serviceModel "github.com/sarastee/auth/internal/model"
	"github.com/sarastee/auth/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, request *user_v1.DeleteRequest) (*emptypb.Empty, error) {
	err := i.userService.Delete(ctx, serviceModel.UserID(request.Id))
	if err != nil {
		return nil, status.Error(codes.Internal, "Delete error")
	}

	return &emptypb.Empty{}, nil
}
