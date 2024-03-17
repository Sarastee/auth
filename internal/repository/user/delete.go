package user

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/sarastee/platform_common/pkg/db"
)

// Delete ...
func (r *Repo) Delete(ctx context.Context, userID int64) error {
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
