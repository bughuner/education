package question

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model/model_view"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func GetQuestionApi(c *gin.Context) {
	var req *model_view.QuestionReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON question_req failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if req.PageNo == 0 {
		req.PageNo = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 100
	}
	questionList, err := GetQuestionList(c, req)
	if err != nil {
		log.Printf("getQuestionList failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, questionList)
}

func GetQuestionList(c *gin.Context, req *model_view.QuestionReq) (*model_view.QuestionResp, error) {
	questionDb := database.Query.Question
	sql := questionDb.WithContext(c)
	if req.ID != "" {
		sql = sql.Where(questionDb.ID.Eq(req.ID))
	}
	if req.Level != 0 {
		sql = sql.Where(questionDb.Level.Eq(req.Level))
	}
	if req.Type != 0 {
		sql = sql.Where(questionDb.Type.Eq(req.Type))
	}
	if req.Damage != 0 {
		sql = sql.Where(questionDb.Damage.Eq(req.Damage))
	}
	total, err := sql.Count()
	if err != nil {
		log.Printf("questionDb count failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("questionDb count failed, err:%v", err)
	}
	questionList, err := sql.Offset(req.PageNo - 1).Limit(req.PageSize).Find()
	if err != nil {
		log.Printf("questionDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("questionDb query failed, err:%v", err)
	}
	questionResult := make([]*model_view.QuestionResult, len(questionList))
	for i, item := range questionList {
		questionSelect, err := getQuestionSelectByQuestionId(c, nil, item.ID)
		if err != nil {
			log.Printf("getQuestionSelectByQuestionId failed, err:%v\n", err)
			continue
		}
		questionAnswer, err := getQuestionAnswerByQuestionId(c, nil, item.ID)
		if err != nil {
			log.Printf("getQuestionAnswerByQuestionId failed, err:%v\n", err)
			continue
		}
		questionResult[i] = &model_view.QuestionResult{
			Question:       item,
			QuestionSelect: questionSelect,
			QuestionAnswer: questionAnswer,
		}
	}
	res := &model_view.QuestionResp{
		Total: total,
		Data:  questionResult,
	}
	return res, nil
}
