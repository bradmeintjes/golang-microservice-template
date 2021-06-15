package postgres

import (
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"

	"sample-microservice-v2/internal/config"
)

func NewConn(c config.Postgres) (*sql.DB, error) {
	return sql.Open("pgx", c.DBName)
}
