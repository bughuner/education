package doc_question

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/model/model_view"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func DeleteDocQuestionApi(c *gin.Context) {
	var req model_view.DeleteDocQuestion
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON DeleteDocQuestion failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	err = deleteDocQuestion(c, req.DocQuestionIds)
	if err != nil {
		log.Printf("deleteDocQuestion failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, nil)
}

func deleteDocQuestion(c *gin.Context, ids []string) error {
	docQuestionDb := database.Query.DocQuestion
	_, err := docQuestionDb.WithContext(c).Where(docQuestionDb.ID.In(ids...)).Delete(&model.DocQuestion{})
	if err != nil {
		log.Printf("docQuestionDb delete failed, err:%v", err)
		return util.BuildErrorInfo("docQuestionDb delete failed, err:%v", err)
	}
	return nil
}
