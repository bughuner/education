package question

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func GetQuestionApi(c *gin.Context) {
	var req model.Question
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON question failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	questionList, err := getQuestionList(c, &req)
	if err != nil {
		log.Printf("getQuestionList failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, questionList)
}

func getQuestionList(c *gin.Context, req *model.Question)([]*model.Question, error) {
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
	questionList, err := sql.Find()
	if err != nil {
		log.Printf("questionDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("questionDb query failed, err:%v", err)
	}
	return questionList, nil
}