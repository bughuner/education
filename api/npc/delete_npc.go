package npc

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

func DeleteNpcApi(c *gin.Context) {
	var req model_view.DeleteNpcReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON DeleteNpcReq failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	err = deleteNpc(c, req.NpcIds)
	if err != nil {
		log.Printf("deleteNpc failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, nil)
}

func deleteNpc(c *gin.Context, ids []string) error {
	npcDb := database.Query.Npc
	_, err := npcDb.WithContext(c).Where(npcDb.ID.In(ids...)).Delete(&model.Npc{})
	if err != nil {
		log.Printf("npcDb delete failed, err:%v\n", err)
		return util.BuildErrorInfo("npcDb delete failed, err:%v", err)
	}
	return nil
}
