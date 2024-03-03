package user

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/sarastee/auth/internal/client/db"
	serviceModel "github.com/sarastee/auth/internal/model"
)

func (r *Repo) Delete(ctx context.Context, userID serviceModel.UserID) error {
	builderDelete := r.sq.Delete(tableDB).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{idDB: userID})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
