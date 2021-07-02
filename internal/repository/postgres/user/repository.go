package user

import (
	"database/sql"

	"webservice-template/internal/domain/user"
)

type Repository struct {
	conn *sql.DB
}

func NewRepository(conn *sql.DB) Repository {
	return Repository{
		conn: conn,
	}
}

func (r Repository) Store(user user.User) error {
	u := User{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	_ = u
	return nil
}
