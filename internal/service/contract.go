package service

import (
	"context"

	"github.com/sarastee/auth/internal/model"
)

// UserService ...
type UserService interface {
	Create(context.Context, *model.UserCreate) (int64, error)
	Update(context.Context, *model.UserUpdate) error
	Get(context.Context, int64) (*model.User, error)
	Delete(context.Context, int64) error
}
