package api

import (
	"api-gatware/cmd/globar"
	__ "api-gatware/cmd/proto"
	"api-gatware/handler/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Shopadd(c *gin.Context) {
	var req request.ShopReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    false,
		})
		return
	}
	shopadd, err := globar.Clients.Shopadd(c, &__.ShopaddReq{
		UserId:  int64(c.GetUint("UserId")),
		Title:   req.Title,
		Info:    req.Info,
		Price:   float32(req.Price),
		ShopNum: req.ShopNum,
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
		"data":    shopadd,
	})
	return
}

func ShopUpdate(c *gin.Context) {
	var req request.ShopReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    false,
		})
		return
	}
	update, err := globar.Clients.ShopUpdate(c, &__.ShopUpdateReq{
		UserId:  int64(c.GetUint("userId")),
		Title:   req.Title,
		Info:    req.Info,
		Price:   float32(req.Price),
		ShopNum: req.ShopNum,
		Id:      req.Id,
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
		"data":    update,
	})
	return
}

func ShopShow(c *gin.Context) {
	var req request.ShopReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    false,
		})
		return
	}
	show, err := globar.Clients.ShopShow(c, &__.ShopShowReq{
		UserId: int64(c.GetUint("UserId")),
		Id:     req.Id,
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
		"message": "详情查看成功",
		"data":    show,
	})
	return
}

func Shoplist(c *gin.Context) {
	var req request.ShopReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    false,
		})
		return
	}
	shoplist, err := globar.Clients.Shoplist(c, &__.ShoplistReq{
		UserId: int64(c.GetUint("UserId")),
		Page:   req.Page,
		Size:   req.Size,
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
		"message": "查看成功",
		"data":    shoplist,
	})
	return
}

func ShopDelete(c *gin.Context) {
	var req request.ShopReq

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    false,
		})
		return
	}
	shopDelete, err := globar.Clients.ShopDelete(c, &__.ShopDeleteReq{
		UserId: int64(c.GetUint("UserId")),
		Id:     req.Id,
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
		"data":    shopDelete,
	})
	return
}
