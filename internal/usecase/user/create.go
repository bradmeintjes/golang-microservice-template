package user

import (
	"sample-microservice-v2/internal/domain/user"
)

type Usecase struct {
	userService user.Service
}

func NewUsecase(userService user.Service) Usecase {
	return Usecase{
		userService: userService,
	}
}

func (u Usecase) Create(user user.User) error {
	return u.userService.Create(user)
}
