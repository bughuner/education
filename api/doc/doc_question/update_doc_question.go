package doc_question

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

func UpdateDocQuestionApi(c *gin.Context) {
	var req model.DocQuestion
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON doc_question failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkUpdateParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	monster, err := updateDocQuestion(c, req.ID, req.DocID, req.QuestionID)
	if err != nil {
		log.Printf("updateDocQuestion failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, monster)
}

func checkUpdateParam(docQuestion *model.DocQuestion) error {
	if docQuestion.ID == "" {
		return util.BuildErrorInfo("ID为空")
	}
	if docQuestion.QuestionID == "" {
		return util.BuildErrorInfo("问题ID为空")
	}
	return nil
}

func updateDocQuestion(c *gin.Context, id, docId, questionId string) (*model.DocQuestion, error) {
	docQuestionDb := database.Query.DocQuestion
	oldDocQuestion, err := docQuestionDb.WithContext(c).Where(docQuestionDb.ID.Eq(id)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("docQuestionDb query failed, id:%v, err:%v\n", id, err)
		return nil, util.BuildErrorInfo("docQuestionDb query failed, err:%v", err)
	}
	if oldDocQuestion == nil {
		return nil, util.BuildErrorInfo("文章问题不存在")
	}
	if docId != "" {
		oldDocQuestion.DocID = docId
	}
	if questionId != "" {
		oldDocQuestion.QuestionID = questionId
	}
	err = docQuestionDb.WithContext(c).Save(oldDocQuestion)
	if err != nil {
		log.Printf("docQuestionDb save failed, docQuestion:%v, err:%v", oldDocQuestion, err)
		return nil, util.BuildErrorInfo("docQuestionDb save failed, err:%v", err)
	}
	return oldDocQuestion, nil
}
