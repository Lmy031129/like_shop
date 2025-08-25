package api

import (
	"api-gatware/cmd/globar"
	__ "api-gatware/cmd/proto"
	"api-gatware/handler/request"
	"api-gatware/until"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req request.UserReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "验证失败",
			"data":    false,
		})
		return
	}
	register, err := globar.Clients.Register(c, &__.RegisterReq{
		Mobile:   req.Mobile,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "添加失败",
			"data":    false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "添加成功",
		"data":    register,
	})
	return
}

func Login(c *gin.Context) {
	var req request.UserReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "验证失败",
			"data":    false,
		})
		return
	}
	login, err := globar.Clients.Login(c, &__.LoginReq{
		Mobile:   req.Mobile,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "登录失败",
			"data":    false,
		})
		return
	}
	j := until.NewJWT()
	claims := until.CustomClaims{ID: uint(login.Id)}
	token, _ := j.CreateToken(claims)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "登录成功",
		"data":    token,
	})
}

func UserShow(c *gin.Context) {
	var req request.UserReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "验证失败",
			"data":    false,
		})
		return
	}
	show, err := globar.Clients.UserShow(c, &__.UserShowReq{
		Mobile: req.Mobile,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "登录失败",
			"data":    false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "查看成功",
		"data":    show,
	})
	return
}
