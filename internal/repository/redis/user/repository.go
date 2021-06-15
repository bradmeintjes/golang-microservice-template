package user

import (
	"sample-microservice-v2/internal/domain/user"
)

type Repository struct{}

func NewRepository() Repository {
	return Repository{}
}

func (r Repository) Cache(user user.User) error {
	return nil
}
