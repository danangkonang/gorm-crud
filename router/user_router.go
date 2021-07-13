package router

import (
	"github.com/danangkonang/gin-rest-api/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/", controller.GetUsers)
}
