package repository

import (
	"context"

	serviceModel "github.com/sarastee/auth/internal/model"
)

// UserRepository ...
type UserRepository interface {
	Create(context.Context, *serviceModel.UserCreate) (int64, error)
	Update(context.Context, *serviceModel.UserUpdate) error
	Get(context.Context, int64) (*serviceModel.User, error)
	Delete(context.Context, int64) error
}
