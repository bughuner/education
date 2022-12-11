package npc

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

func UpdateNpcApi(c *gin.Context) {
	var req model.Npc
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON npc failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkUpdateParam(req.ID); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	npc, err := updateNpc(c, &req)
	if err != nil {
		log.Printf("updateNpc failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, npc)
}

func checkUpdateParam(id string) error {
	if id == "" {
		return util.BuildErrorInfo("id 为空")
	}
	return nil
}

func updateNpc(c *gin.Context, req *model.Npc) (*model.Npc, error) {
	npcDb := database.Query.Npc
	npc, err := npcDb.WithContext(c).Where(npcDb.ID.Eq(req.ID)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("npcDb query failed, err:%v", err)
		return nil, err
	}
	if npc == nil {
		return nil, util.BuildErrorInfo("未找到NPC")
	}
	if req.MapID != "" {
		npc.MapID = req.MapID
	}
	if req.Type != 0 {
		npc.Type = req.Type
	}
	if req.Name != "" {
		npc.Name = req.Name
	}
	if req.Image != "" {
		npc.Image = req.Image
	}
	if req.Introduction != "" {
		npc.Introduction = req.Introduction
	}
	if npc.Sculpt != "" {
		npc.Sculpt = req.Sculpt
	}
	npc.Ordinate = req.Ordinate
	npc.Horizon = req.Horizon
	err = npcDb.WithContext(c).Save(npc)
	if err != nil {
		log.Printf("npcDb save failed, err:%v", err)
		return nil, err
	}
	return npc, nil
}
