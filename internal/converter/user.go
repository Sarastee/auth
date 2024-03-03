package converter

import (
	"github.com/sarastee/auth/internal/common"
	serviceModel "github.com/sarastee/auth/internal/model"
	"github.com/sarastee/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToServiceUserCreateFromCreateRequest(request *user_v1.CreateRequest) *serviceModel.UserCreate {
	return &serviceModel.UserCreate{
		Name:     serviceModel.UserName(request.Name),
		Email:    serviceModel.UserEmail(request.Email),
		Password: serviceModel.Password(request.Password),
		Role:     ToServiceRoleFromRole(request.Role),
	}
}

func ToServiceUserUpdateFromUpdateRequest(request *user_v1.UpdateRequest) *serviceModel.UserUpdate {
	var (
		name  *serviceModel.UserName
		email *serviceModel.UserEmail
		role  *serviceModel.UserRole
	)

	if request.Name == nil {
		name = nil
	} else {
		name = common.Pointer[serviceModel.UserName](serviceModel.UserName(*request.Name))
	}

	if request.Email == nil {
		email = nil
	} else {
		email = common.Pointer[serviceModel.UserEmail](serviceModel.UserEmail(*request.Email))
	}

	if request.Role == nil {
		role = nil
	} else {
		role = common.Pointer[serviceModel.UserRole](ToServiceRoleFromRole(*request.Role))
	}

	return &serviceModel.UserUpdate{
		ID:    serviceModel.UserID(request.Id),
		Name:  name,
		Email: email,
		Role:  role,
	}

}

func ToServiceRoleFromRole(role user_v1.Role) serviceModel.UserRole {
	roleName := user_v1.Role_name[int32(role)]

	return serviceModel.UserRole(roleName)
}

func ToRoleFromServiceRole(role serviceModel.UserRole) user_v1.Role {
	resultRole := user_v1.Role_value[string(role)]

	return user_v1.Role(resultRole)
}

func ToGetResponseFromServiceUser(user *serviceModel.User) *user_v1.GetResponse {
	return &user_v1.GetResponse{
		Id:        int64(user.ID),
		Name:      string(user.Name),
		Email:     string(user.Email),
		Role:      ToRoleFromServiceRole(user.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdateAt:  timestamppb.New(user.UpdatedAt),
	}
}
