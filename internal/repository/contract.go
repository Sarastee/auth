package repository

import (
	"context"

	serviceModel "github.com/sarastee/auth/internal/model"
)

//go:generate ../../bin/mockery --output ./mocks  --inpackage-suffix --all

// UserRepository interface for CRUD user repository
type UserRepository interface {
	Create(context.Context, *serviceModel.UserCreate) (int64, error)
	Update(context.Context, *serviceModel.UserUpdate) error
	Get(context.Context, int64) (*serviceModel.User, error)
	Delete(context.Context, int64) error
}
