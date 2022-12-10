package monster_question

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

func GetMonsterQuestionApi(c *gin.Context) {
	var req model_view.GetMonsterQuestionReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON GetMonsterQuestionReq failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkGetParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	res, err := getMonsterQuestion(c, &req)
	if err != nil {
		log.Printf("getMonsterInfo failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, res)
}

func checkGetParam(req *model_view.GetMonsterQuestionReq) error {
	if req.PageNo == 0 {
		req.PageNo = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 100
	}
	return nil
}

func getMonsterQuestion(c *gin.Context, req *model_view.GetMonsterQuestionReq) (*model_view.GetMonsterQuestionResp, error) {
	monsterQuestionDb := database.Query.MonsterQuestion
	sql := monsterQuestionDb.WithContext(c)
	if req.ID != "" {
		sql = sql.Where(monsterQuestionDb.ID.Eq(req.ID))
	}
	if req.MonsterID != "" {
		sql = sql.Where(monsterQuestionDb.MonsterID.Eq(req.MonsterID))
	}
	total, err := sql.Count()
	if err != nil {
		log.Printf("monsterQuestionDb count failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("monsterQuestionDb count failed, err:%v", err)
	}
	monsterQuestionList, err := sql.Offset(req.PageNo - 1).Limit(req.PageSize).Find()
	if err != nil {
		log.Printf("monsterQuestionDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("monsterQuestionDb query failed, err:%v\n", err)
	}
	getMonsterQuestionRes := make([]*model_view.GetMonsterQuestionRes, len(monsterQuestionList))
	for i, item := range monsterQuestionList {
		question, err := question.GetQuestionList(c, &model_view.QuestionReq{ID: item.QuestionID, PageNo: 1, PageSize: 1})
		if err != nil || question == nil || len(question.Data) == 0 {
			log.Printf("GetQuestionList failed, err:%v", err)
			continue
		}
		getMonsterQuestionRes[i] = &model_view.GetMonsterQuestionRes{
			MonsterQuestion: item,
			Data:            question.Data[0],
		}
	}
	res := &model_view.GetMonsterQuestionResp{
		Total: total,
		Data:  getMonsterQuestionRes,
	}

	return res, nil
}
