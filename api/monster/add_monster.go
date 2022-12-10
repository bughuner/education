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

func AddMonsterApi(c *gin.Context) {
	var req model.Monster
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON monster failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkAddParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	npc, err := addMonster(c, &req)
	if err != nil {
		log.Printf("addDoc failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, npc)
}

func checkAddParam(req *model.Monster) error {
	if req.Name == "" {
		return util.BuildErrorInfo("名称为空")
	}
	if req.Type == 0 {
		return util.BuildErrorInfo("类型为空")
	}
	if req.Level <= 0 {
		return util.BuildErrorInfo("等级<=0")
	}
	if req.Hp <= 0 {
		return util.BuildErrorInfo("血量<=0")
	}
	if req.Coin <= 0 {
		return util.BuildErrorInfo("金币<=0")
	}
	if req.ExperienceValue <= 0 {
		return util.BuildErrorInfo("经验值<=0")
	}
	return nil
}

func addMonster(c *gin.Context, monster *model.Monster) (*model.Monster, error) {
	monsterDb := database.Query.Monster
	id := util.GetUUID()
	monster.ID = id
	err := monsterDb.WithContext(c).Create(monster)
	if err != nil {
		log.Printf("monsterDb create failed, err:%v", err)
		return nil, err
	}
	return monster, nil
}
