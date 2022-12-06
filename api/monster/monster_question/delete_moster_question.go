package monster_question

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

func DeleteMonsterQuestionApi(c *gin.Context) {
	var req model_view.DeleteMonsterQuestion
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON monster_question_ids failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	err = deleteMonsterQuestion(c, req.MonsterQuestionIds)
	if err != nil {
		log.Printf("deleteMonsterQuestion failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, nil)
}

func deleteMonsterQuestion(c *gin.Context, ids []string) error {
	monsterQuestionDb := database.Query.MonsterQuestion
	_, err := monsterQuestionDb.WithContext(c).Where(monsterQuestionDb.ID.In(ids...)).Delete(&model.MonsterQuestion{})
	if err != nil {
		log.Printf("monsterQuestionDb delete failed, err:%v", err)
		return util.BuildErrorInfo("monsterQuestionDb delete failed, err:%v", err)
	}
	return nil
}
