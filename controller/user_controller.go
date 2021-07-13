package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	// name := c.Param("name")
	// action := c.Param("action")
	// message := name + " is " + action
	// c.String(http.StatusOK, message)
	c.JSON(http.StatusOK, "users")
}
