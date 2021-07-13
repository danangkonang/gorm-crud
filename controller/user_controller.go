package controller

import (
	"github.com/danangkonang/gin-rest-api/entity"
	"github.com/danangkonang/gin-rest-api/service"
)

type UserController interface {
	FindAllUser() []entity.User
}

type controller struct {
	service service.UserService
}

func New(userService service.UserService) UserController {
	return &controller{
		service: userService,
	}
}
func (c *controller) FindAllUser() []entity.User {
	return c.service.FindAllUser()
}
