package monster_question

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

func UpdateMonsterQuestionApi(c *gin.Context) {
	var req model.MonsterQuestion
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON monster_question failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkUpdateParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	monster, err := updateMonsterQuestion(c, req.ID, req.MonsterID, req.QuestionID)
	if err != nil {
		log.Printf("updateMonster failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, monster)
}

func checkUpdateParam(monsterQuestion *model.MonsterQuestion) error {
	if monsterQuestion.ID == "" {
		return util.BuildErrorInfo("ID为空")
	}
	if monsterQuestion.QuestionID == "" {
		return util.BuildErrorInfo("问题ID为空")
	}
	return nil
}

func updateMonsterQuestion(c *gin.Context, id, monsterId, questionId string) (*model.MonsterQuestion, error) {
	monsterQuestionDb := database.Query.MonsterQuestion
	oldMonsterQuestion, err := monsterQuestionDb.WithContext(c).Where(monsterQuestionDb.ID.Eq(id)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("monsterQuestionDb query failed, id:%v, err:%v\n", id, err)
		return nil, util.BuildErrorInfo("monsterQuestionDb query failed, err:%v", err)
	}
	if oldMonsterQuestion == nil {
		return nil, util.BuildErrorInfo("怪物问题不存在")
	}
	if monsterId != "" {
		oldMonsterQuestion.MonsterID = monsterId
	}
	if questionId != "" {
		oldMonsterQuestion.QuestionID = questionId
	}
	err = monsterQuestionDb.WithContext(c).Save(oldMonsterQuestion)
	if err != nil {
		log.Printf("monsterQuestionDb save failed, monsterQuestion:%v, err:%v", oldMonsterQuestion, err)
		return nil, util.BuildErrorInfo("monsterQuestionDb save failed, err:%v", err)
	}
	return oldMonsterQuestion, nil
}
