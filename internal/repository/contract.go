package repository

import (
	"context"

	serviceModel "github.com/sarastee/auth/internal/model"
)

type UserRepository interface {
	Create(context.Context, *serviceModel.UserCreate) (serviceModel.UserID, error)
	Update(context.Context, *serviceModel.UserUpdate) error
	Get(context.Context, serviceModel.UserID) (*serviceModel.User, error)
	Delete(context.Context, serviceModel.UserID) error
}
