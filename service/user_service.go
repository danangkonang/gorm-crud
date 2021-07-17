package service

import (
	"github.com/danangkonang/gin-rest-api/entity"
	"github.com/danangkonang/gin-rest-api/repository"
)

type UserService interface {
	FindAllUser() []entity.User
	SaveUser(users entity.User) entity.User
}

type userService struct {
	repository repository.Repository
}

func New(repository repository.Repository) UserService {
	return &userService{
		repository: repository,
	}
}

func (u *userService) FindAllUser() []entity.User {
	return u.repository.FindUsers()
}

func (u *userService) SaveUser(users entity.User) entity.User {
	u.repository.SaveUser(users)
	return users
}
