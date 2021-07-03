package postgres

import (
	"context"
	"fmt"

	"webservice-template/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewConn(ctx context.Context, c config.Postgres) (*sqlx.DB, error) {
	connUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.User, c.Password, c.Host, c.Port, c.DBName, c.SSL)
	conn, err := sqlx.ConnectContext(ctx, "postgres", connUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return conn, nil
}
