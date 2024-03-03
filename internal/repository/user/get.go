package user

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/sarastee/auth/internal/client/db"
	serviceModel "github.com/sarastee/auth/internal/model"
	"github.com/sarastee/auth/internal/repository"
	repoConverter "github.com/sarastee/auth/internal/repository/user/converter"
	repoModel "github.com/sarastee/auth/internal/repository/user/model"
)

func (r *Repo) Get(ctx context.Context, UserID serviceModel.UserID) (*serviceModel.User, error) {
	builderSelect := r.sq.Select(idDB, emailDB, nameDB, roleDB, createdAtDB, updatedAtDB).
		From(tableDB).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{idDB: UserID}).
		Limit(1)

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[repoModel.User])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repository.ErrUserNotFound
		}
		return nil, err
	}

	return repoConverter.ToServiceUserFromRepoUser(&out), nil
}
