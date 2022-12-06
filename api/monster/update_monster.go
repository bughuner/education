package monster

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func UpdateMonsterApi(c *gin.Context) {
	var req model.Monster
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON monster failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkUpdateParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	oldMonster, err := getMonsterInfo(c, req.ID)
	if err != nil {
		log.Printf("getMonsterInfo failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	monster, err := updateMonster(c, oldMonster, &req)
	if err != nil {
		log.Printf("updateMonster failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, monster)
}

func checkUpdateParam(monster *model.Monster) error {
	if monster.ID == "" {
		return util.BuildErrorInfo("ID为空")
	}
	return nil
}

func updateMonster(c *gin.Context, monster *model.Monster, req *model.Monster) (*model.Monster, error) {
	if req.Name != "" {
		monster.Name = req.Name
	}
	if req.Coin != 0 {
		monster.Coin = req.Coin
	}
	if req.ExperienceValue != 0 {
		monster.ExperienceValue = req.ExperienceValue
	}
	if req.Level != 0 {
		monster.Level = req.Level
	}
	if req.Type != 0 {
		monster.Type = req.Type
	}
	if req.Image != "" {
		monster.Image = req.Image
	}
	monster.Hp = req.Hp

	monsterDb := database.Query.Monster
	err := monsterDb.WithContext(c).Save(monster)
	if err != nil {
		log.Printf("monsterDb save failed, monster:%v, req:%v, err:%v", monster, req, err)
		return nil, util.BuildErrorInfo("monsterDb save failed, err:%v", err)
	}
	return monster, nil
}
