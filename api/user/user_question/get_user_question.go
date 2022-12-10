package user_question

import (
	"education/api/question"
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model/model_view"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func GetUserQuestionApi(c *gin.Context) {
	var req model_view.GetUserQuestionReq
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
	userQuestion, err := getUserQuestion(c, &req)
	if err != nil {
		log.Printf("updateUserQuestion failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, userQuestion)
}

func checkGetUserQuestionParams(req *model_view.GetUserQuestionReq) error {
	if req.PageNo == 0 {
		req.PageNo = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 100
	}
	return nil
}

func getUserQuestion(c *gin.Context, req *model_view.GetUserQuestionReq) (*model_view.GetUserQuestionResp, error) {
	userQuestionDb := database.Query.UserQuestion
	sql := userQuestionDb.WithContext(c)
	if req.ID != "" {
		sql = sql.Where(userQuestionDb.ID.Eq(req.ID))
	}
	if req.UserID != "" {
		sql = sql.Where(userQuestionDb.UserID.Eq(req.UserID))
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
		if err != nil || question == nil || len(question.Data) == 0 {
			log.Printf("GetQuestionList failed, err:%v", err)
			continue
		}
		userQuestionRes[i] = &model_view.UserQuestionRes{
			UserQuestion: item,
			Data:         question.Data[0],
		}
	}
	res := &model_view.GetUserQuestionResp{
		Total: total,
		Data:  userQuestionRes,
	}
	return res, nil
}
