package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/sarastee/auth/internal/model"
	"github.com/sarastee/auth/internal/repository/mocks"
	userService "github.com/sarastee/auth/internal/service/user"
	"github.com/stretchr/testify/require"
)

func userCreate() *model.UserCreate {
	return &model.UserCreate{
		Name:     "test",
		Email:    "test@mail.com",
		Password: "password",
		Role:     "role",
	}
}

func TestCreate_SuccessCreateUser(t *testing.T) {
	ctx := context.Background()

	user := userCreate()
	var userID int64 = 1

	userRepoMock := mocks.NewUserRepository(t)
	userRepoMock.On("Create", ctx, user).Return(userID, nil).Once()

	service := userService.NewService(userRepoMock)

	resultUserID, err := service.Create(ctx, user)

	require.NoError(t, err)
	require.NotEmpty(t, resultUserID)
	require.Equal(t, userID, resultUserID)
}

func TestCreate_FailCreateUser(t *testing.T) {
	ctx := context.Background()

	user := userCreate()
	repoErr := errors.New("some error")

	userRepoMock := mocks.NewUserRepository(t)
	userRepoMock.On("Create", ctx, user).Return(int64(0), repoErr).Once()

	service := userService.NewService(userRepoMock)

	_, err := service.Create(ctx, user)

	require.Error(t, err)
}
