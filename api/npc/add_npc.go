package npc

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func AddNpcApi(c *gin.Context) {
	var req model.Npc
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON npc failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkAddParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	npc, err := addNpc(c, &req)
	if err != nil {
		log.Printf("addDoc failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, npc)
}

func checkAddParam(req *model.Npc) error {
	if req.Name == "" {
		return util.BuildErrorInfo("名称为空")
	}
	if req.Type == 0 {
		return util.BuildErrorInfo("类型为空")
	}
	if req.MapID == "" {
		return util.BuildErrorInfo("地图ID为空")
	}
	return nil
}

func addNpc(c *gin.Context, npc *model.Npc) (*model.Npc, error) {
	npcDb := database.Query.Npc
	id := util.GetUUID()
	npc.ID = id
	err := npcDb.WithContext(c).Create(npc)
	if err != nil {
		log.Printf("npcDb create failed, err:%v", err)
		return nil, err
	}
	return npc, nil
}
