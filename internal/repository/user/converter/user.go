package converter

import (
	"github.com/sarastee/auth/internal/common"
	serviceModel "github.com/sarastee/auth/internal/model"
	repoModel "github.com/sarastee/auth/internal/repository/user/model"
)

func ToRepoUserFromServiceUser(user *serviceModel.User) *repoModel.User {
	return &repoModel.User{
		ID:        int64(user.ID),
		Name:      string(user.Name),
		Email:     string(user.Email),
		Role:      string(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToRepoUserCreateFromServiceUserCreate(user *serviceModel.UserCreate) *repoModel.UserCreate {
	return &repoModel.UserCreate{
		Name:     string(user.Name),
		Email:    string(user.Email),
		Password: string(user.Password),
		Role:     string(user.Role),
	}
}

func ToRepoUserUpdateFromServiceUserUpdate(user *serviceModel.UserUpdate) *repoModel.UserUpdate {
	var (
		name  *string
		email *string
		role  *string
	)

	if user.Name == nil {
		name = nil
	} else {
		name = common.Pointer[string](string(*user.Name))
	}

	if user.Email == nil {
		email = nil
	} else {
		email = common.Pointer[string](string(*user.Email))
	}

	if user.Role == nil {
		role = nil
	} else {
		role = common.Pointer[string](string(*user.Role))
	}

	return &repoModel.UserUpdate{
		ID:    int64(user.ID),
		Name:  name,
		Email: email,
		Role:  role,
	}
}

func ToServiceUserFromRepoUser(user *repoModel.User) *serviceModel.User {
	return &serviceModel.User{
		ID:        serviceModel.UserID(user.ID),
		Name:      serviceModel.UserName(user.Name),
		Email:     serviceModel.UserEmail(user.Email),
		Role:      serviceModel.UserRole(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
