package tests

import (
	"context"
	"errors"
	"testing"

	userImpl "github.com/sarastee/auth/internal/api/user"
	"github.com/sarastee/auth/internal/repository/mocks"
	userService "github.com/sarastee/auth/internal/service/user"
	"github.com/stretchr/testify/require"
)

func TestCreate_SuccessUpdateUser(t *testing.T) {
	ctx := context.Background()

	var userID int64 = 1
	request := userUpdateRequest(userID)

	userRepoMock := mocks.NewUserRepository(t)
	userRepoMock.On("Update", ctx, userUpdateRepo(userID)).Return(nil).Once()

	service := userService.NewService(userRepoMock)

	impl := userImpl.NewImplementation(service)

	_, err := impl.Update(ctx, request)

	require.NoError(t, err)
}

func TestCreate_FailUpdateUser(t *testing.T) {
	ctx := context.TODO()

	var userID int64 = 1
	request := userUpdateRequest(userID)
	someErr := errors.New("some error")

	userRepoMock := mocks.NewUserRepository(t)
	userRepoMock.On("Update", ctx, userUpdateRepo(userID)).Return(someErr).Once()

	service := userService.NewService(userRepoMock)

	impl := userImpl.NewImplementation(service)

	_, err := impl.Update(ctx, request)

	require.Error(t, err)
}
