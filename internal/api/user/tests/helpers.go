package tests

import (
	"time"

	"github.com/sarastee/auth/internal/common"
	"github.com/sarastee/auth/internal/model"
	"github.com/sarastee/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func userCreateRequest() *user_v1.CreateRequest {
	return &user_v1.CreateRequest{
		Name:            "test",
		Email:           "test",
		Password:        "test",
		PasswordConfirm: "test",
		Role:            1,
	}
}

func userCreateForRepo() *model.UserCreate {
	return &model.UserCreate{
		Name:     "test",
		Email:    "test",
		Password: "test",
		Role:     "USER",
	}
}

func userResponseForGet(userID int64, userDate time.Time) *user_v1.GetResponse {
	return &user_v1.GetResponse{
		Id:        userID,
		Name:      "test",
		Email:     "test",
		Role:      1,
		CreatedAt: timestamppb.New(userDate),
		UpdatedAt: timestamppb.New(userDate),
	}
}

func userFromRepo(userID int64, userDate time.Time) *model.User {
	return &model.User{
		ID:        userID,
		Name:      "test",
		Email:     "test",
		Role:      "USER",
		CreatedAt: userDate,
		UpdatedAt: userDate,
	}
}

func userUpdateRequest(userID int64) *user_v1.UpdateRequest {
	return &user_v1.UpdateRequest{
		Id:    userID,
		Name:  common.Pointer[string]("test"),
		Email: common.Pointer[string]("test"),
		Role:  common.Pointer[user_v1.Role](1),
	}
}

func userUpdateRepo(userID int64) *model.UserUpdate {
	return &model.UserUpdate{
		ID:    userID,
		Name:  common.Pointer[string]("test"),
		Email: common.Pointer[string]("test"),
		Role:  common.Pointer[model.UserRole]("USER"),
	}
}
