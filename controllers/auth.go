package controllers

import (
	"chuns.cn/oj/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginController(c *gin.Context) {
	var loginUser models.User
	c.ShouldBind(&loginUser)
	fmt.Println(loginUser)
	user, err := models.GetUserByUsername(loginUser.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "2001",
			"msg":  "User does not exist.",
			"data": "",
		})
		return
	}
	if user.Password != models.Encrypt(loginUser.Password) {
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

func RegisterController(c *gin.Context) {
	var user models.User
	c.ShouldBind(&user)
	fmt.Println(user)

	_, err1 := models.GetUserByUsername(user.Username)
	if err1 == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "2001",
			"msg":  "Username existed.",
			"data": "",
		})
		return
	}

	_, err2 := models.GetUserByEmail(user.Email)
	if err2 == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "2002",
			"msg":  "Email has been registered.",
			"data": "",
		})
		return
	}

	err := user.Create()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "2003",
			"msg":  "create user failed.",
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
