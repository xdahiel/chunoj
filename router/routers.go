package router

import (
	"chuns.cn/oj/controllers"
	"chuns.cn/oj/databases"
	"chuns.cn/oj/setting"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	databases.InitMySQL()
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	r.Use(cors.Default())

	r.POST("/login", controllers.LoginController)

	return r
}
