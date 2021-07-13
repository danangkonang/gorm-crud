package api

import (
	"github.com/danangkonang/gin-rest-api/controller"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
	userController controller.UserController
}

func NewApi(userConreoller controller.UserController) *UserApi {
	return &UserApi{
		userController: userConreoller,
	}
}

func (api *UserApi) FindAllUser(ctx *gin.Context) {
	ctx.JSON(200, api.userController.FindAllUser())
}
