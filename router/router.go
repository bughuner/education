package router

import (
	."education/api/monster"
	. "education/api/user"
	."education/api/question"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	// 用户相关
	user := router.Group("user")
	{
		user.POST("/login",LoginApi)
		user.POST("/add_user",AddUserApi)
		user.POST("/update_user",UpdateUserApi)
	}
	monster := router.Group("monster")
	{
		monster.POST("/get_monster", GetMonsterApi)
		monster.POST("/update_monster",UpdateMonsterApi)
	}
	question := router.Group("question")
	{
		question.POST("/get_question", GetQuestionApi)
	}
}
