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

func AddUserQuestionApi(c *gin.Context) {
	var req model.UserQuestion
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON user_question failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err := checkAddUserQuestionParams(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	userQuestion, err := addUserQuestion(c, req.UserID, req.QuestionID)
	if err != nil {
		log.Printf("addUserQuestion failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, userQuestion)
}

func checkAddUserQuestionParams(req *model.UserQuestion) error {
	if req.UserID == "" {
		return util.BuildErrorInfo("用户ID为空")
	}
	if req.QuestionID == "" {
		return util.BuildErrorInfo("问题ID为空")
	}
	return nil
}

func addUserQuestion(c *gin.Context, userId, questionId string) (*model.UserQuestion, error) {
	id := util.GetUUID()
	userQuestion := &model.UserQuestion{
		ID:         id,
		UserID:     userId,
		QuestionID: questionId,
	}
	userQuestionDb := database.Query.UserQuestion
	err := userQuestionDb.WithContext(c).Create(userQuestion)
	if err != nil {
		log.Printf("userQuestionDb create failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("userQuestionDb create failed, err:%v", err)
	}
	return userQuestion, nil
}
