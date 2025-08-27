package router

import (
	"api-gatware/handler/api"
	"api-gatware/until"

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
	video := r.Group("/video")
	{
		video.GET("videolist", api.VideoList)
		video.Use(until.JWTAuth())
		video.POST("videoadd", api.Videoadd)
		video.POST("videoupdate", api.Videoupdate)
		video.POST("videodelete", api.VideoDelete)
	}

	return r
}
