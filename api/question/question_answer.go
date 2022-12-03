package question

import (
	"education/database"
	"education/model"
	"education/query"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func addQuestionAnswer(c *gin.Context, query *query.Query, questionId string, answer []string) error {
	if query == nil {
		query = database.Query
	}
	questionSelectMap, err := getQuestionSelectByContent(c, query, questionId, answer)
	if err != nil {
		return util.BuildErrorInfo("getQuestionSelect failed, err:%v", err)
	}
	questionAnswerDb := query.QuestionAnswer
	questionAnswerList := make([]*model.QuestionAnswer, len(answer))
	for i, item := range answer {
		id := util.GetUUID()
		qs, ok := questionSelectMap[item]
		if !ok {
			return util.BuildErrorInfo("答案不存在")
		}
		questionAnswerList[i] = &model.QuestionAnswer{
			ID:         id,
			QuestionID: questionId,
			Answer:     qs.ID,
		}
	}
	err = questionAnswerDb.WithContext(c).Create(questionAnswerList...)
	if err != nil {
		log.Printf("questionAnswerDb create failed, err:%v\n", err)
		return util.BuildErrorInfo("questionAnswerDb create failed, err:%v", err)
	}
	return nil
}

func getQuestionAnswerByQuestionId(c *gin.Context, query *query.Query, questionId string) ([]*model.QuestionAnswer, error) {
	if query == nil {
		query = database.Query
	}
	questionAnswerDb := query.QuestionAnswer
	questionAnswerList, err := questionAnswerDb.WithContext(c).Where(questionAnswerDb.QuestionID.Eq(questionId)).Find()
	if err != nil {
		log.Printf("questionSelectDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("questionSelectDb query failed, err:%v", err)
	}
	return questionAnswerList, nil
}
