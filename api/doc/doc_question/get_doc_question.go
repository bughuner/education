package doc_question

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

func GetDocQuestionApi(c *gin.Context) {
	var req model_view.GetDocQuestionReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON GetDocQuestionReq failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkGetParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	res, err := getDocQuestion(c, &req)
	if err != nil {
		log.Printf("getDocQuestion failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, res)
}

func checkGetParam(req *model_view.GetDocQuestionReq) error {
	if req.PageNo == 0 {
		req.PageNo = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 100
	}
	return nil
}

func getDocQuestion(c *gin.Context, req *model_view.GetDocQuestionReq) (*model_view.GetDocQuestionResp, error) {
	docQuestionDb := database.Query.DocQuestion
	sql := docQuestionDb.WithContext(c)
	if req.ID != "" {
		sql = sql.Where(docQuestionDb.ID.Eq(req.ID))
	}
	if req.DocId != "" {
		sql = sql.Where(docQuestionDb.DocID.Eq(req.DocId))
	}
	total, err := sql.Count()
	if err != nil {
		log.Printf("docQuestionDb count failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("docQuestionDb count failed, err:%v", err)
	}
	docQuestionList, err := sql.Offset(req.PageNo - 1).Limit(req.PageSize).Find()
	if err != nil {
		log.Printf("docQuestionDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("docQuestionDb query failed, err:%v\n", err)
	}
	getDocQuestionRes := make([]*model_view.GetDocQuestionRes, len(docQuestionList))
	for i, item := range docQuestionList {
		question, err := question.GetQuestionList(c, &model_view.QuestionReq{ID: item.QuestionID, PageNo: 1, PageSize: 1})
		if err != nil || len(question.Data) == 0 {
			log.Printf("GetQuestionList failed, err:%v", err)
			continue
		}
		getDocQuestionRes[i] = &model_view.GetDocQuestionRes{
			DocQuestion: item,
			Data:        question.Data[0],
		}
	}
	res := &model_view.GetDocQuestionResp{
		Total: total,
		Data:  getDocQuestionRes,
	}

	return res, nil
}
