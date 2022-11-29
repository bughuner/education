package router

import (
	. "education/api"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	// 用户相关
	user := router.Group("user")
	{
		user.GET("/get_user",GetUserApi)
		user.POST("/add_user",AddUserApi)
	}
}
