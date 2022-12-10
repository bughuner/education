package monster

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

func DeleteMonsterApi(c *gin.Context) {
	var req model_view.DeleteMonsterReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON DeleteMonsterReq failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	err = deleteMonster(c, req.MonsterIds)
	if err != nil {
		log.Printf("deleteMonster failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, nil)
}

func deleteMonster(c *gin.Context, ids []string) error {
	monsterDb := database.Query.Monster
	_, err := monsterDb.WithContext(c).Where(monsterDb.ID.In(ids...)).Delete(&model.Monster{})
	if err != nil {
		log.Printf("monsterDb delete failed, err:%v\n", err)
		return util.BuildErrorInfo("monsterDb delete failed, err:%v", err)
	}
	return nil
}
