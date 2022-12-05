package user_question

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func UpdateUserQuestionApi(c *gin.Context) {
	var req model.UserQuestion
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON user_question failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err := checkUpdateUserQuestionParams(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	userQuestion, err := updateUserQuestion(c, req.ID, req.UserID, req.QuestionID)
	if err != nil {
		log.Printf("updateUserQuestion failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, userQuestion)
}

func checkUpdateUserQuestionParams(req *model.UserQuestion) error {
	if req.ID == "" {
		return util.BuildErrorInfo("ID为空")
	}
	if req.UserID == "" {
		return util.BuildErrorInfo("用户ID为空")
	}
	if req.QuestionID == "" {
		return util.BuildErrorInfo("问题ID为空")
	}
	return nil
}

func updateUserQuestion(c *gin.Context, id, userId, questionId string) (*model.UserQuestion, error) {
	userQuestionDb := database.Query.UserQuestion
	userQuestion, err := userQuestionDb.WithContext(c).Where(userQuestionDb.ID.Eq(id)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("userQuestionDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("userQuestionDb query failed, err:%v", err)
	}
	if userQuestion == nil {
		return nil, util.BuildErrorInfo("用户问题不存在")
	}
	if userId != "" {
		userQuestion.UserID = userId
	}
	if questionId != "" {
		userQuestion.QuestionID = questionId
	}
	err = userQuestionDb.WithContext(c).Save(userQuestion)
	if err != nil {
		log.Printf("userQuestionDb save failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("userQuestionDb save failed, err:%v", err)
	}
	return userQuestion, nil
}
