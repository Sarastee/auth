package user

import (
	"github.com/Masterminds/squirrel"
	"github.com/sarastee/auth/internal/repository"
	"github.com/sarastee/platform_common/pkg/db"
)

const (
	tableDB = "users"

	idDB                   = "id"
	nameDB                 = "name"
	emailDB                = "email"
	passwordDB             = "password"
	roleDB                 = "role"
	createdAtDB            = "created_at"
	updatedAtDB            = "updated_at"
	userEmailKeyConstraint = "user_email_key"
)

var _ repository.UserRepository = (*Repo)(nil)

// Repo ...
type Repo struct {
	db db.Client
	sq squirrel.StatementBuilderType
}

// NewRepository ...
func NewRepository(client db.Client) *Repo {
	return &Repo{
		db: client,
		sq: squirrel.StatementBuilder,
	}
}
