package user

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	storager Storager
	cacher   Cacher
}

func NewService(storager Storager, cacher Cacher) Service {
	return Service{
		storager: storager,
		cacher:   cacher,
	}
}

func (s Service) Create(user User) error {
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()

	return s.storager.Store(user)
}
