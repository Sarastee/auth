package service

import (
	"context"

	"github.com/sarastee/auth/internal/model"
)

type UserService interface {
	Create(context.Context, *model.UserCreate) (model.UserID, error)
	Update(context.Context, *model.UserUpdate) error
	Get(context.Context, model.UserID) (*model.User, error)
	Delete(context.Context, model.UserID) error
}
