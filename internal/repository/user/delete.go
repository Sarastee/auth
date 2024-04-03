package user

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/sarastee/platform_common/pkg/db"
)

// Delete ...
func (r *Repo) Delete(ctx context.Context, userID int64) error {
	queryFormat := `
	DELETE 
	FROM 
	    %s 
	WHERE 
	    %s = @%s
	`

	query := fmt.Sprintf(
		queryFormat,
		usersTable,
		idColumn, idColumn)

	q := db.Query{
		Name:     "user_repository.Delete",
		QueryRaw: query,
	}

	args := pgx.NamedArgs{
		idColumn: userID,
	}

	_, err := r.db.DB().ExecContext(ctx, q, args)
	if err != nil {
		return err
	}

	return nil
}
