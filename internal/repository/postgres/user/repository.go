package user

import (
	"log"
	"webservice-template/internal/domain/user"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	conn *sqlx.DB
}

func NewRepository(conn *sqlx.DB) Repository {
	return Repository{
		conn: conn,
	}
}

func (r Repository) All() ([]user.User, error) {
	users := []User{}
	err := r.conn.Select(&users, "select * from users")

	if err == nil {
		return mapToDomainUsers(users), nil
	}

	return []user.User{}, err
}

func (r Repository) Store(user user.User) error {
	u := User{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	tx := r.conn.MustBegin()

	res, err := tx.NamedExec("insert into users (name) values (:name)", &u)
	log.Printf("%+v", res)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func mapToDomainUsers(users []User) []user.User {
	mapped := make([]user.User, len(users))
	for i, v := range users {
		mapped[i] = mapToDomainUser(v)
	}
	return mapped
}

func mapToDomainUser(usr User) user.User {
	return user.User{
		ID:        usr.ID,
		Name:      usr.Name,
		CreatedAt: usr.CreatedAt,
		UpdatedAt: usr.UpdatedAt,
	}
}
