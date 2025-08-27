package api

import (
	"api-gatware/cmd/globar"
	__ "api-gatware/cmd/proto"
	"api-gatware/handler/request"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Videoadd(c *gin.Context) {
	var req request.VideoReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    false,
		})
		return
	}
	videoAdd, err := globar.Clients.Videoadd(c, &__.VideoaddReq{
		Userid:   int64(c.GetUint("userId")),
		Title:    req.Title,
		VideoUrl: req.VideoUrl,
		Info:     req.Info,
		ImageUrl: req.ImageUrl,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "添加成功",
		"data":    videoAdd,
	})
	return

}

func Videoupdate(c *gin.Context) {
	var req request.VideoReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    false,
		})
		return
	}
	videoUpdate, err := globar.Clients.Videoupdate(c, &__.VideoupdateReq{
		Id:     int64(req.ID),
		Userid: int64(c.GetUint("userId")),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "修改成功",
		"data":    videoUpdate,
	})
	return
}

func VideoDelete(c *gin.Context) {
	var req request.VideoReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    false,
		})
		return
	}
	videoDelete, err := globar.Clients.Videodelete(c, &__.VideodeleteReq{
		Id:     int64(req.ID),
		Userid: int64(c.GetUint("userId")),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "删除成功",
		"data":    videoDelete,
	})
	return

}

func VideoList(c *gin.Context) {
	var req request.VideoReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    false,
		})
		return
	}
	videoList, err := globar.Clients.Videolist(c, &__.VideolistReq{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    false,
		})
		return
	}
	fmt.Println("算数", 1024*4)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "查看成功",
		"data":    videoList,
	})
	return
}
