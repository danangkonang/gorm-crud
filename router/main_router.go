package router

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func Run() {
	routes()
	router.Run(":5000")
}

func routes() {
	v1 := router.Group("/v1")
	UserRoutes(v1)
}
