package main

import "github.com/danangkonang/gin-rest-api/router"

func main() {
	router.Run()

	// router.Run()
	// server := gin.Default()
	// rest := api.NewApi(controller.New(service.New(repository.NewRepository())))
	// r := server.Group("/v1")
	// r.GET("/users", rest.FindAllUser)
	// server.Run(":8000")
}
