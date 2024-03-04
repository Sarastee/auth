package converter

import (
	"github.com/sarastee/auth/internal/common"
	serviceModel "github.com/sarastee/auth/internal/model"
	"github.com/sarastee/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToServiceUserCreateFromCreateRequest ...
func ToServiceUserCreateFromCreateRequest(request *user_v1.CreateRequest) *serviceModel.UserCreate {
	return &serviceModel.UserCreate{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Role:     ToServiceRoleFromRole(request.Role),
	}
}

// ToServiceUserUpdateFromUpdateRequest ...
func ToServiceUserUpdateFromUpdateRequest(request *user_v1.UpdateRequest) *serviceModel.UserUpdate {
	var (
		name  *string
		email *string
		role  *serviceModel.UserRole
	)

	if request.Name == nil {
		name = nil
	} else {
		name = common.Pointer[string](*request.Name)
	}

	if request.Email == nil {
		email = nil
	} else {
		email = common.Pointer[string](*request.Email)
	}

	if request.Role == nil {
		role = nil
	} else {
		role = common.Pointer[serviceModel.UserRole](ToServiceRoleFromRole(*request.Role))
	}

	return &serviceModel.UserUpdate{
		ID:    request.Id,
		Name:  name,
		Email: email,
		Role:  role,
	}

}

// ToServiceRoleFromRole ...
func ToServiceRoleFromRole(role user_v1.Role) serviceModel.UserRole {
	roleName := user_v1.Role_name[int32(role)]

	return serviceModel.UserRole(roleName)
}

// ToRoleFromServiceRole ...
func ToRoleFromServiceRole(role serviceModel.UserRole) user_v1.Role {
	resultRole := user_v1.Role_value[string(role)]

	return user_v1.Role(resultRole)
}

// ToGetResponseFromServiceUser ...
func ToGetResponseFromServiceUser(user *serviceModel.User) *user_v1.GetResponse {
	return &user_v1.GetResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      ToRoleFromServiceRole(user.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdateAt:  timestamppb.New(user.UpdatedAt),
	}
}
