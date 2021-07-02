package postgres

import (
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"

	"webservice-template/internal/config"
)

func NewConn(c config.Postgres) (*sql.DB, error) {
	return sql.Open("pgx", c.DBName)
}
