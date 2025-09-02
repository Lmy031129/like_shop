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
	shop := r.Group("/shop")
	{
		shop.Use(until.JWTAuth())
		shop.POST("shopadd", api.Shopadd)
		shop.POST("shopUpdate", api.ShopUpdate)
		shop.POST("shopShow", api.ShopShow)
		shop.POST("shopDelete", api.ShopDelete)
		shop.GET("shoplist", api.Shoplist)
	}

	return r
}
