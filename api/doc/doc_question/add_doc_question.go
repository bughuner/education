package doc_question

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func AddDocQuestionApi(c *gin.Context) {
	var req model.DocQuestion
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON doc_question failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkAddParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	docQuestion, err := addDocQuestion(c, req.DocID, req.QuestionID)
	if err != nil {
		log.Printf("addDocQuestion failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, docQuestion)
}

func checkAddParam(req *model.DocQuestion) error {
	if req.QuestionID == "" {
		return util.BuildErrorInfo("问题ID为空")
	}
	if req.DocID == "" {
		return util.BuildErrorInfo("文章ID为空")
	}
	return nil
}

func addDocQuestion(c *gin.Context, docId, questionId string) (*model.DocQuestion, error) {
	id := util.GetUUID()
	docQuestion := &model.DocQuestion{
		ID:         id,
		DocID:      docId,
		QuestionID: questionId,
	}
	docQuestionDb := database.Query.DocQuestion
	err := docQuestionDb.WithContext(c).Create(docQuestion)
	if err != nil {
		log.Printf("docQuestionDb create failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("docQuestionDb create failed, err:%v", err)
	}
	return docQuestion, nil
}
