package router

import (
	"github.com/danangkonang/gin-rest-api/controller"
	"github.com/danangkonang/gin-rest-api/repository"
	"github.com/danangkonang/gin-rest-api/service"
	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) {
	rest := controller.New(service.New(repository.NewRepository()))
	rg.GET("/users", rest.FindAllUser)
	rg.POST("/user", rest.SaveUser)
	rg.PUT("/user", rest.UpdateUser)
	rg.DELETE("/user/:id", rest.DeleteUser)
	rg.GET("/user/:id", rest.DetailUser)
}
