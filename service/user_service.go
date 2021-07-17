package service

import (
	"github.com/danangkonang/gin-rest-api/entity"
	"github.com/danangkonang/gin-rest-api/repository"
)

type UserService interface {
	FindAllUser() []entity.User
	SaveUser(users entity.User) entity.User
	DeleteUser(users entity.User) entity.User
	UpdateUser(users entity.User) entity.User
	DetailUser(users entity.User) entity.User
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

func (u *userService) DeleteUser(users entity.User) entity.User {
	u.repository.DeleteUser(&users)
	return users
}

func (u *userService) UpdateUser(users entity.User) entity.User {
	u.repository.UpdateUser(&users)
	return users
}

func (u *userService) DetailUser(users entity.User) entity.User {
	u.repository.FindUserById(&users)
	return users
}
