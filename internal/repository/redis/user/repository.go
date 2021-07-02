package user

import (
	"webservice-template/internal/domain/user"
)

type Repository struct{}

func NewRepository() Repository {
	return Repository{}
}

func (r Repository) Cache(user user.User) error {
	return nil
}
