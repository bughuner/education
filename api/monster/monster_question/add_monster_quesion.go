package monster_question

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func AddMonsterQuestionApi(c *gin.Context) {
	var req model.MonsterQuestion
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON monster_question failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkAddParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	monsterQuestion, err := addMonsterQuestion(c, req.MonsterID, req.QuestionID)
	if err != nil {
		log.Printf("addMonsterQuestion failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, monsterQuestion)
}

func checkAddParam(req *model.MonsterQuestion) error {
	if req.QuestionID == "" {
		return util.BuildErrorInfo("问题ID为空")
	}
	if req.MonsterID == "" {
		return util.BuildErrorInfo("怪物ID为空")
	}
	return nil
}

func addMonsterQuestion(c *gin.Context, monsterId, questionId string) (*model.MonsterQuestion, error) {
	id := util.GetUUID()
	monsterQuestion := &model.MonsterQuestion{
		ID:         id,
		MonsterID:  monsterId,
		QuestionID: questionId,
	}
	monsterQuestionDb := database.Query.MonsterQuestion
	err := monsterQuestionDb.WithContext(c).Create(monsterQuestion)
	if err != nil {
		log.Printf("monsterQuestionDb create failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("monsterQuestionDb create failed, err:%v", err)
	}
	return monsterQuestion, nil
}
