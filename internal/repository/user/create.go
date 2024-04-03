package user

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
	serviceModel "github.com/sarastee/auth/internal/model"
	"github.com/sarastee/auth/internal/repository"
	"github.com/sarastee/platform_common/pkg/db"
)

// Create ...
func (r *Repo) Create(ctx context.Context, in *serviceModel.UserCreate) (int64, error) {
	currentTime := time.Now()

	queryFormat := `
	INSERT INTO 
	    %s (%s, %s, %s, %s, %s, %s) 
	VALUES 
		(@%s, @%s, @%s, @%s, @%s, @%s) 
	RETURNING id
	`

	query := fmt.Sprintf(
		queryFormat,
		usersTable,
		nameColumn, emailColumn, roleColumn, passwordColumn, createdAtColumn, updatedAtColumn,
		nameColumn, emailColumn, roleColumn, passwordColumn, createdAtColumn, updatedAtColumn,
	)

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	args := pgx.NamedArgs{
		nameColumn:      in.Name,
		emailColumn:     in.Email,
		roleColumn:      in.Role,
		passwordColumn:  in.Password,
		createdAtColumn: currentTime,
		updatedAtColumn: currentTime,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var UserID int64

	UserID, err = pgx.CollectOneRow(rows, pgx.RowTo[int64])
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.ConstraintName == userEmailKeyConstraint {
			return 0, repository.ErrEmailIsTaken
		}
	}

	return UserID, nil
}
