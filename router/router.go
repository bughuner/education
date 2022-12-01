package router

import (
	. "education/api/exchange_gift"
	. "education/api/monster"
	. "education/api/npc"
	. "education/api/question"
	. "education/api/shop"
	. "education/api/user"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	// 用户相关
	user := router.Group("user")
	{
		user.POST("/login",LoginApi)
		user.POST("/add_user",AddUserApi)
		user.POST("/update_user",UpdateUserApi)
		user.POST("/add_user_task", AddUserTaskApi)
		user.GET("/get_user_task", GetUserTaskApi)
	}
	monster := router.Group("monster")
	{
		monster.GET("/get_monster", GetMonsterApi)
		monster.POST("/update_monster",UpdateMonsterApi)
	}
	question := router.Group("question")
	{
		question.POST("/get_question", GetQuestionApi)
	}
	npc := router.Group("npc")
	{
		npc.GET("/get_npc_task", GetNPCTaskApi)
	}
	shop := router.Group("shop")
	{
		shop.GET("/get_gift_count", GetGiftCountApi)
	}
	exchange := router.Group("/exchange_gift")
	{
		exchange.POST("/buy_gift",ExchangeGiftApi)
	}

}
