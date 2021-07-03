package user

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	storage Storage
	cacher  Cacher
}

func NewService(storage Storage, cacher Cacher) Service {
	return Service{
		storage: storage,
		cacher:  cacher,
	}
}

func (s Service) Create(user User) error {
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()

	return s.storage.Store(user)
}

func (s Service) All() ([]User, error) {
	return s.storage.All()
}
