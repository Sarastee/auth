package user

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/sarastee/auth/internal/client/db"
	serviceModel "github.com/sarastee/auth/internal/model"
	"github.com/sarastee/auth/internal/repository/user/converter"
)

// Update ...
func (r *Repo) Update(ctx context.Context, user *serviceModel.UserUpdate) error {
	repoUser := converter.ToRepoUserUpdateFromServiceUserUpdate(user)

	builderUpdate := r.sq.Update(tableDB).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{idDB: repoUser.ID})

	if repoUser.Name != nil {
		builderUpdate = builderUpdate.Set(nameDB, repoUser.Name)
	}
	if repoUser.Email != nil {
		builderUpdate = builderUpdate.Set(emailDB, repoUser.Email)
	}
	if repoUser.Role != nil {
		builderUpdate = builderUpdate.Set(roleDB, repoUser.Role)
	}

	builderUpdate = builderUpdate.Set(updatedAtDB, time.Now())

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
