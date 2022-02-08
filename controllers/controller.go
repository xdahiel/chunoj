package controllers

import (
	"chuns.cn/oj/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": "",
	})
}

func LoginController(c *gin.Context) {
	param := c.Params
	user, err := models.GetUserByUsername(param.ByName("Username"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "2001",
			"msg":  "User does not exist.",
			"data": "",
		})
		return
	}
	if user.Password != models.Encrypt(param.ByName("password")) {
		c.JSON(http.StatusOK, gin.H{
			"code": "2002",
			"msg":  "Incorrect password",
			"data": "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "2000",
		"msg":  "success",
		"data": user,
	})
}
