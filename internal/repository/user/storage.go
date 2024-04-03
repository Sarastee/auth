package user

import (
	"github.com/sarastee/auth/internal/repository"
	"github.com/sarastee/platform_common/pkg/db"
)

const (
	usersTable = "users"

	idColumn               = "id"
	nameColumn             = "name"
	emailColumn            = "email"
	passwordColumn         = "password"
	roleColumn             = "role"
	createdAtColumn        = "created_at"
	updatedAtColumn        = "updated_at"
	userEmailKeyConstraint = "user_email_key"
)

var _ repository.UserRepository = (*Repo)(nil)

// Repo ...
type Repo struct {
	db db.Client
}

// NewRepository ...
func NewRepository(client db.Client) *Repo {
	return &Repo{
		db: client,
	}
}
