package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	serviceModel "github.com/sarastee/auth/internal/model"
	"github.com/sarastee/auth/internal/repository"
	repoConverter "github.com/sarastee/auth/internal/repository/user/converter"
	repoModel "github.com/sarastee/auth/internal/repository/user/model"
	"github.com/sarastee/platform_common/pkg/db"
)

// Get ...
func (r *Repo) Get(ctx context.Context, UserID int64) (*serviceModel.User, error) {
	queryFormat := `
	SELECT
	    %s, %s, %s, %s, %s, %s 
	FROM 
	    %s 
	WHERE 
	    %s = @%s
	`

	query := fmt.Sprintf(
		queryFormat,
		idColumn, emailColumn, nameColumn, roleColumn, createdAtColumn, updatedAtColumn,
		usersTable,
		idColumn, idColumn)

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	args := pgx.NamedArgs{
		idColumn: UserID,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args)
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
