package user_question

import (
	"education/api/question"
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/model/model_view"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func GetUserQuestionApi(c *gin.Context) {
	var req model.UserQuestion
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON user_question failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err := checkGetUserQuestionParams(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	userQuestion, err := getUserQuestion(c, req.ID, req.UserID, req.QuestionID)
	if err != nil {
		log.Printf("updateUserQuestion failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, userQuestion)
}

func checkGetUserQuestionParams(req *model.UserQuestion) error {
	return nil
}

func getUserQuestion(c *gin.Context, id, userId, questionId string) (*model_view.UserQuestionResp, error) {
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
	total, err := sql.Count()
	if err != nil {
		log.Printf("userQuestionDb count failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("userQuestionDb query failed, err:%v", err)
	}
	userQuestionList, err := sql.Find()
	if err != nil {
		log.Printf("userQuestionDb count failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("userQuestionDb query failed, err:%v", err)
	}
	userQuestionRes := make([]*model_view.UserQuestionRes, len(userQuestionList))
	for i, item := range userQuestionList {
		question, err := question.GetQuestionList(c, &model_view.QuestionReq{ID: item.QuestionID, PageNo: 1, PageSize: 1})
		if err != nil || len(question.Data) > 0 {
			log.Printf("GetQuestionList failed, err:%v", err)
			continue
		}
		userQuestionRes[i] = &model_view.UserQuestionRes{
			UserQuestion: item,
			Data:         question.Data[0],
		}
	}
	res := &model_view.UserQuestionResp{
		Total: total,
		Data:  userQuestionRes,
	}
	return res, nil
}
