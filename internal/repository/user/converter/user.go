package converter

import (
	"github.com/sarastee/auth/internal/common"
	serviceModel "github.com/sarastee/auth/internal/model"
	repoModel "github.com/sarastee/auth/internal/repository/user/model"
)

// ToRepoUserUpdateFromServiceUserUpdate ...
func ToRepoUserUpdateFromServiceUserUpdate(user *serviceModel.UserUpdate) *repoModel.UserUpdate {
	var (
		name  *string
		email *string
		role  *string
	)

	if user.Name == nil {
		name = nil
	} else {
		name = common.Pointer[string](*user.Name)
	}

	if user.Email == nil {
		email = nil
	} else {
		email = common.Pointer[string](*user.Email)
	}

	if user.Role == nil {
		role = nil
	} else {
		role = common.Pointer[string](string(*user.Role))
	}

	return &repoModel.UserUpdate{
		ID:    user.ID,
		Name:  name,
		Email: email,
		Role:  role,
	}
}

// ToServiceUserFromRepoUser ...
func ToServiceUserFromRepoUser(user *repoModel.User) *serviceModel.User {
	return &serviceModel.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      serviceModel.UserRole(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
