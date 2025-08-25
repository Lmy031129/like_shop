package router

import (
	"api-gatware/handler/api"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	user := r.Group("/user")
	{
		user.POST("/register", api.Register)
		user.POST("/login", api.Login)
		user.POST("/userShow", api.UserShow)

	}

	return r
}
