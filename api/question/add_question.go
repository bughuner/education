package question

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/model/model_view"
	"education/query"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func AddQuestionApi(c *gin.Context) {
	var req model_view.QuestionReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON question_all failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err := checkAddParam(&req); err != nil {
		log.Printf("checkAddParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	questionList, err := addQuestion(c, &req)
	if err != nil {
		log.Printf("addQuestion failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, questionList)
}

func checkAddParam(req *model_view.QuestionReq) error {
	if req.Damage <= 0 {
		return util.BuildErrorInfo("伤害小于0")
	}
	if req.Level <= 0 {
		return util.BuildErrorInfo("等级小于0")
	}
	if req.Question == "" {
		return util.BuildErrorInfo("问题为空")
	}
	if req.Type <= 0 && req.Type >= 4 {
		return util.BuildErrorInfo("未知的题目类型")
	}
	if len(req.Select) <= 1 {
		return util.BuildErrorInfo("选项个数必须大于1")
	}
	if len(req.Answer) <= 0 {
		return util.BuildErrorInfo("答案个数必须大于0")
	}
	for _, item := range req.Answer {
		flag := false
		for _, s := range req.Select {
			if item == s {
				flag = true
			}
		}
		if !flag {
			return util.BuildErrorInfo("答案不在选项中,答案%v", item)
		}
	}
	return nil
}

func addQuestion(c *gin.Context, question *model_view.QuestionReq) (*model_view.QuestionReq, error) {
	id := util.GetUUID()
	questionEntity := &model.Question{
		ID:       id,
		Level:    question.Level,
		Type:     question.Type,
		Question: question.Question,
		Damage:   question.Damage,
	}
	err := database.Query.Transaction(func(tx *query.Query) error {
		err := tx.Question.WithContext(c).Create(questionEntity)
		if err != nil {
			return err
		}
		err = addQuestionSelect(c, tx, questionEntity.ID, question.Select)
		if err != nil {
			return err
		}
		err = addQuestionAnswer(c, tx, questionEntity.ID, question.Answer)
		return nil
	})
	if err != nil {
		log.Printf("question execute transaction failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("question execute transaction failed, err:%v", err)
	}
	return nil, nil
}
