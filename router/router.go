package router

import (
	. "education/api/doc"
	. "education/api/doc/doc_question"
	. "education/api/exchange_gift"
	. "education/api/form"
	. "education/api/gift"
	. "education/api/monster"
	. "education/api/monster/monster_question"
	. "education/api/npc"
	. "education/api/npc/npc_task"
	. "education/api/plat"
	. "education/api/question"
	. "education/api/shop"
	. "education/api/task"
	. "education/api/user"
	"education/api/user/user_question"
	"education/api/user/user_task"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	// 用户相关
	user := router.Group("user")
	{
		user.POST("/login", LoginApi)
		user.POST("/get_user", GetUserApi)
		user.POST("/add_user", AddUserApi)
		user.POST("/update_user", UpdateUserApi)
		user.POST("/add_user_task", user_task.AddUserTaskApi)
		user.POST("/get_user_task", user_task.GetUserTaskApi)
		user.POST("/update_user_task", user_task.UpdateUserTaskApi)
		user.POST("/add_user_question", user_question.AddUserQuestionApi)
		user.POST("/get_user_question", user_question.GetUserQuestionApi)
		user.POST("/update_user_question", user_question.UpdateUserQuestionApi)
		user.POST("/delete_user_question", user_question.DeleteUserQuestionApi)
	}
	monster := router.Group("monster")
	{
		monster.POST("/add_monster", AddMonsterApi)
		monster.POST("/get_monster", GetMonsterApi)
		monster.POST("/update_monster", UpdateMonsterApi)
		monster.POST("/delete_monster", DeleteMonsterApi)
		monster.POST("/get_monster_question", GetMonsterQuestionApi)
		monster.POST("/add_monster_question", AddMonsterQuestionApi)
		monster.POST("/update_monster_question", UpdateMonsterQuestionApi)
		monster.POST("/delete_monster_question", DeleteMonsterQuestionApi)
	}
	question := router.Group("question")
	{
		question.POST("/get_question", GetQuestionApi)
		question.POST("/add_question", AddQuestionApi)
	}
	npc := router.Group("npc")
	{
		npc.POST("/get_npc", GetNpcApi)
		npc.POST("/add_npc", AddNpcApi)
		npc.POST("/update_npc", UpdateNpcApi)
		npc.POST("/delete_npc", DeleteNpcApi)
		npc.POST("/get_npc_task", GetNPCTaskApi)
		npc.POST("/add_npc_task", AddNpcTaskApi)
		npc.POST("/update_npc_task", UpdateNpcApi)
		npc.POST("/delete_npc_task", DeleteNpcTaskApi)
	}
	shop := router.Group("shop")
	{
		shop.POST("/get_gift_count", GetGiftCountApi)
		shop.POST("/add_gift_count", AddGiftCountApi)
		shop.POST("/update_gift_count", UpdateGiftCountApi)
	}
	exchange := router.Group("/exchange_gift")
	{
		exchange.POST("/buy_gift", ExchangeGiftApi)
		exchange.POST("/get_exchange_gift", GetExchangeGiftApi)
		exchange.POST("/update_exchange_gift", UpdateExchangeGiftApi)
	}
	doc := router.Group("/doc")
	{
		doc.POST("/add_doc", AddDocApi)
		doc.POST("/get_doc", GetDocApi)
		doc.POST("/update_doc", UpdateDocApi)
		doc.POST("/delete_doc", DeleteDocApi)
		doc.POST("/get_doc_question", GetDocQuestionApi)
		doc.POST("/add_doc_question", AddDocQuestionApi)
		doc.POST("/update_doc_question", UpdateDocQuestionApi)
		doc.POST("/delete_doc_question", DeleteDocQuestionApi)
	}
	task := router.Group("/task")
	{
		task.POST("/add_task", AddTaskApi)
		task.POST("/get_task", GetTaskApi)
		task.POST("update_task", UpdateTaskApi)
		task.POST("/delete_task", DeleteTaskApi)
	}
	form := router.Group("/form")
	{
		form.POST("/add_form", AddFormApi)
		form.POST("/get_form", GetFormApi)
		form.POST("/update_form", UpdateFormApi)
		form.POST("/delete_form", DeleteFormApi)
	}
	plat := router.Group("/plat")
	{
		plat.POST("/get_plat", GetPlatApi)
		plat.POST("/add_plat", AddPlatApi)
		plat.POST("/update_plat", UpdatePlatApi)
		plat.POST("/delete_plat", DeletePlatApi)
	}
	gift := router.Group("/gift")
	{
		gift.POST("/get_gift", GetGiftApi)
		gift.POST("/add_gift", AddGiftApi)
		gift.POST("/update_gift", UpdateGiftApi)
		gift.POST("/delete_gift", DeleteGiftApi)
	}
}
