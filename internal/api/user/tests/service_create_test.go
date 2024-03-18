package tests

import (
	"context"
	"errors"
	"testing"

	userImpl "github.com/sarastee/auth/internal/api/user"
	"github.com/sarastee/auth/internal/repository"
	"github.com/sarastee/auth/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sarastee/auth/internal/repository/mocks"
	userService "github.com/sarastee/auth/internal/service/user"
	"github.com/stretchr/testify/require"
)

func TestCreate_SuccessCreateUser(t *testing.T) {
	ctx := context.Background()

	var userID int64 = 1
	request := userCreateRequest()

	userRepoMock := mocks.NewUserRepository(t)
	userRepoMock.On("Create", ctx, userCreateForRepo()).Return(userID, nil).Once()

	service := userService.NewService(userRepoMock)

	impl := userImpl.NewImplementation(service)

	resultUserID, err := impl.Create(ctx, request)

	require.NoError(t, err)
	require.NotEmpty(t, resultUserID)
}

func TestCreate_FailCreateUser(t *testing.T) {
	ctx := context.Background()

	request := userCreateRequest()
	repoUserForCreate := userCreateForRepo()

	errorRepoMockGenerator := func(err error) *mocks.UserRepository {
		userRepoMock := mocks.NewUserRepository(t)
		userRepoMock.
			On("Create", ctx, repoUserForCreate).
			Return(int64(0), err).
			Once()

		return userRepoMock
	}

	tests := []struct {
		name              string
		createRequest     *user_v1.CreateRequest
		internalError     error
		expectedErrorCode codes.Code
	}{
		{
			name:              "fail create because email is taken",
			createRequest:     request,
			internalError:     repository.ErrEmailIsTaken,
			expectedErrorCode: codes.AlreadyExists,
		},
		{
			name:              "fail create",
			createRequest:     request,
			internalError:     errors.New("some error"),
			expectedErrorCode: codes.Internal,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repoMock := errorRepoMockGenerator(test.internalError)
			service := userService.NewService(repoMock)
			impl := userImpl.NewImplementation(service)

			_, err := impl.Create(ctx, test.createRequest)

			require.Error(t, err)
			if e, ok := status.FromError(err); ok {
				require.Equal(t, e.Code(), test.expectedErrorCode)
			}
		})
	}
}
