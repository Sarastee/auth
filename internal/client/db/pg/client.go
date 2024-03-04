package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sarastee/auth/internal/client/db"
)

type pgClient struct {
	masterDBC db.DB
}

// New ...
func New(ctx context.Context, dsn string) (db.Client, error) {
	pgxConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("error while creating config for pgxpool: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, fmt.Errorf("fail to connect to DB: %v", err)
	}

	return &pgClient{
		masterDBC: NewDB(pool),
	}, nil
}

func (c *pgClient) DB() db.DB {
	return c.masterDBC
}

func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}
