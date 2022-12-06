package service

import (
	"education/database"
	"education/model"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func GetDocQuestionByDocId(c *gin.Context, docIds []string) ([]*model.DocQuestion, error) {
	docQuestionDb := database.Query.DocQuestion
	docQuestionList, err := docQuestionDb.WithContext(c).Where(docQuestionDb.DocID.In(docIds...)).Find()
	if err != nil {
		log.Printf("docQuestionDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("docQuestionDb query failed, err:%v\n", err)
	}
	return docQuestionList, nil
}
