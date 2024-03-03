package user

import (
	"context"

	"github.com/sarastee/auth/internal/converter"
	"github.com/sarastee/auth/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, request *user_v1.UpdateRequest) (*emptypb.Empty, error) {
	err := i.userService.Update(ctx, converter.ToServiceUserUpdateFromUpdateRequest(request))
	if err != nil {
		return nil, status.Error(codes.Internal, "Update Error")
	}

	return &emptypb.Empty{}, nil
}
