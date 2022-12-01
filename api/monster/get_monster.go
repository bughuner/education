package monster

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

func GetMonsterApi(c *gin.Context) {
	monsterId := c.Query("monster_id")
	if err := checkGetParam(monsterId); err != nil {
		log.Printf("checkParam failed, monsterId:%v, err:%v\n", monsterId, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	monster, err := getMonsterInfo(c, monsterId)
	if err != nil {
		log.Printf("getMonsterInfo failed, monsterId:%v, err:%v\n", monsterId, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, monster)
}

func checkGetParam(monsterId string) error {
	if monsterId == "" {
		return util.BuildErrorInfo("参数错误")
	}
	return nil
}

func getMonsterInfo(c *gin.Context, id string) (*model.Monster, error) {
	monsterDb := database.Query.Monster
	monster, err := monsterDb.WithContext(c).Where(monsterDb.ID.Eq(id)).First()
	if err != nil && err != gorm.ErrRecordNotFound{
		log.Printf("monsterDb query failed, id:%v, err:%v\n", id, err)
		return nil, util.BuildErrorInfo("monsterDb query failed, err:%v", err)
	}
	if monster == nil {
		return nil, util.BuildErrorInfo("怪物不存在")
	}
	return monster, nil
}