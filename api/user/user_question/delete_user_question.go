package user_question

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func DeleteUserQuestionApi(c *gin.Context) {
	var req model.UserQuestion
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON user_question failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err := checkDeleteUserQuestionParams(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	err = deleteUserQuestion(c, req.ID, req.UserID, req.QuestionID)
	if err != nil {
		log.Printf("deleteUserQuestion failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, nil)
}

func checkDeleteUserQuestionParams(req *model.UserQuestion) error {
	return nil
}

func deleteUserQuestion(c *gin.Context, id, userId, questionId string) error {
	userQuestionDb := database.Query.UserQuestion
	sql := userQuestionDb.WithContext(c)
	if id != "" {
		sql = sql.Where(userQuestionDb.ID.Eq(id))
	}
	if userId != "" {
		sql = sql.Where(userQuestionDb.UserID.Eq(userId))
	}
	if questionId != "" {
		sql = sql.Where(userQuestionDb.QuestionID.Eq(questionId))
	}
	_, err := sql.Delete(&model.UserQuestion{})
	if err != nil {
		log.Printf("userQuestionDb delete failed, err:%v\n", err)
		return util.BuildErrorInfo("userQuestionDb delete failed, err:%v", err)
	}
	return nil
}
