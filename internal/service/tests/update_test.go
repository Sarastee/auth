package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/sarastee/auth/internal/common"
	"github.com/sarastee/auth/internal/model"
	"github.com/sarastee/auth/internal/repository/mocks"
	userService "github.com/sarastee/auth/internal/service/user"
	"github.com/stretchr/testify/require"
)

func userForUpdate() *model.UserUpdate {
	return &model.UserUpdate{
		ID:    1,
		Name:  common.Pointer[string]("test"),
		Email: common.Pointer[string]("test@mail.com"),
		Role:  common.Pointer[model.UserRole]("admin"),
	}
}

func TestUpdate_SuccessUpdateUser(t *testing.T) {
	ctx := context.TODO()

	user := userForUpdate()

	userRepoMock := mocks.NewUserRepository(t)
	userRepoMock.On("Update", ctx, user).Return(nil).Once()

	service := userService.NewService(userRepoMock)

	err := service.Update(ctx, user)

	require.NoError(t, err)
}

func TestUpdate_FailUpdateUser(t *testing.T) {
	ctx := context.TODO()

	user := userForUpdate()
	repoErr := errors.New("some error")

	userRepoMock := mocks.NewUserRepository(t)
	userRepoMock.On("Update", ctx, user).Return(repoErr).Once()

	service := userService.NewService(userRepoMock)

	err := service.Update(ctx, user)

	require.Error(t, err)
}
