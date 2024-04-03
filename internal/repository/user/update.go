package user

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	serviceModel "github.com/sarastee/auth/internal/model"
	"github.com/sarastee/auth/internal/repository/user/converter"
	"github.com/sarastee/platform_common/pkg/db"
)

// Update ...
func (r *Repo) Update(ctx context.Context, user *serviceModel.UserUpdate) error {
	repoUser := converter.ToRepoUserUpdateFromServiceUserUpdate(user)

	queryFormat := `
	UPDATE 
	    %s 
	SET 
	    %s 
	WHERE 
	    %s = @%s
	`

	setParams := make([]string, 0)
	setFormat := "%s = @%s"
	namedArgs := make(pgx.NamedArgs)

	if repoUser.Name != nil {
		setParams = append(setParams, fmt.Sprintf(setFormat, nameColumn, nameColumn))
		namedArgs[nameColumn] = repoUser.Name
	}
	if repoUser.Email != nil {
		setParams = append(setParams, fmt.Sprintf(setFormat, emailColumn, emailColumn))
		namedArgs[emailColumn] = repoUser.Email
	}
	if repoUser.Role != nil {
		setParams = append(setParams, fmt.Sprintf(setFormat, roleColumn, roleColumn))
		namedArgs[roleColumn] = repoUser.Role
	}

	isNeededUpdate := len(namedArgs) > 0

	if !isNeededUpdate {
		return nil
	}

	namedArgs[idColumn] = repoUser.ID

	setParams = append(setParams, fmt.Sprintf(setFormat, updatedAtColumn, updatedAtColumn))
	namedArgs[updatedAtColumn] = time.Now()

	setParamsStr := strings.Join(setParams, ", ")

	query := fmt.Sprintf(
		queryFormat,
		usersTable,
		setParamsStr,
		idColumn, idColumn,
	)

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	_, err := r.db.DB().ExecContext(ctx, q, namedArgs)

	return err
}
