package user

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
	"github.com/sarastee/auth/internal/client/db"
	serviceModel "github.com/sarastee/auth/internal/model"
	"github.com/sarastee/auth/internal/repository"
)

func (r *Repo) Create(ctx context.Context, in *serviceModel.UserCreate) (serviceModel.UserID, error) {
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

	var UserID serviceModel.UserID
	UserID, err = pgx.CollectOneRow(rows, pgx.RowTo[serviceModel.UserID])
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.ConstraintName == userEmailKeyConstraint {
			return 0, repository.ErrEmailIsTaken
		}
	}

	return UserID, nil
}
