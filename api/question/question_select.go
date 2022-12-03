package question

import (
	"education/database"
	"education/model"
	"education/query"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func addQuestionSelect(c *gin.Context, query *query.Query, questionId string, content []string) error {
	if query == nil {
		query = database.Query
	}
	questionSelectDb := query.QuestionSelect
	questionSelectList := make([]*model.QuestionSelect, len(content))
	for i, item := range content {
		id := util.GetUUID()
		questionSelectList[i] = &model.QuestionSelect{
			ID:         id,
			QuestionID: questionId,
			Content:    item,
		}
	}
	err := questionSelectDb.WithContext(c).Create(questionSelectList...)
	if err != nil {
		log.Printf("questionSelectDb create failed, err:%v\n", err)
		return util.BuildErrorInfo("questionSelectDb create failed, err:%v", err)
	}
	return nil
}

func getQuestionSelectByContent(c *gin.Context, query *query.Query, questionId string, content []string) (map[string]*model.QuestionSelect, error) {
	if query == nil {
		query = database.Query
	}
	questionSelectDb := query.QuestionSelect
	questionSelectList, err := questionSelectDb.WithContext(c).Where(questionSelectDb.QuestionID.Eq(questionId), questionSelectDb.Content.In(content...)).Find()
	if err != nil {
		log.Printf("questionSelectDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("questionSelectDb query failed, err:%v", err)
	}
	m := make(map[string]*model.QuestionSelect)
	for _, item := range questionSelectList {
		m[item.Content] = item
	}
	return m, nil
}

func getQuestionSelectByQuestionId(c *gin.Context, query *query.Query, questionId string) ([]*model.QuestionSelect, error) {
	if query == nil {
		query = database.Query
	}
	questionSelectDb := query.QuestionSelect
	questionSelectList, err := questionSelectDb.WithContext(c).Where(questionSelectDb.QuestionID.Eq(questionId)).Find()
	if err != nil {
		log.Printf("questionSelectDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("questionSelectDb query failed, err:%v", err)
	}
	return questionSelectList, nil
}
