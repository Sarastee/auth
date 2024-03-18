package user

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
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
	builderInsert := r.sq.Insert(tableDB).
		PlaceholderFormat(squirrel.Dollar).
		Columns(nameDB, emailDB, passwordDB, roleDB, createdAtDB, updatedAtDB).
		Values(in.Name, in.Email, in.Password, in.Role, currentTime, currentTime).
		Suffix(fmt.Sprintf("RETURNING %s", idDB))

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
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
