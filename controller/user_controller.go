package controller

import (
	"github.com/danangkonang/gin-rest-api/entity"
	"github.com/danangkonang/gin-rest-api/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserService
}

func New(userService service.UserService) *UserController {
	return &UserController{
		service: userService,
	}
}

func (c *UserController) FindAllUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status":  200,
		"message": "success",
		"data":    c.service.FindAllUser(),
	})
}

func (c *UserController) SaveUser(ctx *gin.Context) {
	var usr entity.User
	if err := ctx.ShouldBindJSON(&usr); err != nil {
		ctx.JSON(200, gin.H{
			"status":  200,
			"message": err.Error(),
		})
		return
	}
	ctx.BindJSON(&usr)
	ctx.JSON(200, gin.H{
		"status":  200,
		"message": "success",
		"data":    c.service.SaveUser(usr),
	})
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	var usr entity.User
	if err := ctx.ShouldBindJSON(&usr); err != nil {
		ctx.JSON(200, gin.H{
			"status":  200,
			"message": err.Error(),
		})
		return
	}
	ctx.BindJSON(&usr)
	ctx.JSON(200, gin.H{
		"status":  200,
		"message": "success",
		"data":    c.service.UpdateUser(usr),
	})
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	var usr entity.User
	if err := ctx.ShouldBindJSON(&usr); err != nil {
		ctx.JSON(200, gin.H{
			"status":  200,
			"message": err.Error(),
		})
		return
	}
	ctx.BindJSON(&usr)
	ctx.JSON(200, gin.H{
		"status":  200,
		"message": "success",
		"data":    c.service.DeleteUser(usr),
	})
}

func (c *UserController) DetailUser(ctx *gin.Context) {
	var usr entity.User
	if err := ctx.ShouldBindJSON(&usr); err != nil {
		ctx.JSON(200, gin.H{
			"status":  200,
			"message": err.Error(),
		})
		return
	}
	ctx.BindJSON(&usr)
	ctx.JSON(200, gin.H{
		"status":  200,
		"message": "success",
		"data":    c.service.DetailUser(usr),
	})
}
