package tests

import (
	"context"
	"errors"
	"testing"

	userImpl "github.com/sarastee/auth/internal/api/user"
	"github.com/sarastee/auth/internal/repository/mocks"
	userService "github.com/sarastee/auth/internal/service/user"
	"github.com/sarastee/auth/pkg/user_v1"
	"github.com/stretchr/testify/require"
)

func TestCreate_SuccessDeleteUser(t *testing.T) {
	ctx := context.Background()

	var userID int64 = 1
	request := &user_v1.DeleteRequest{Id: userID}

	userRepoMock := mocks.NewUserRepository(t)
	userRepoMock.On("Delete", ctx, userID).Return(nil).Once()

	service := userService.NewService(userRepoMock)

	impl := userImpl.NewImplementation(service)

	_, err := impl.Delete(ctx, request)

	require.NoError(t, err)
}

func TestCreate_FailDeleteUser(t *testing.T) {
	ctx := context.Background()

	var userID int64 = 1
	request := &user_v1.DeleteRequest{Id: userID}
	someErr := errors.New("some error")

	userRepoMock := mocks.NewUserRepository(t)
	userRepoMock.On("Delete", ctx, userID).Return(someErr).Once()

	service := userService.NewService(userRepoMock)

	impl := userImpl.NewImplementation(service)

	_, err := impl.Delete(ctx, request)

	require.Error(t, err)
}
