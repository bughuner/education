package npc_task

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

func DeleteNpcTaskApi(c *gin.Context) {
	var req model_view.DeleteNpcReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON DeleteNpcReq failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	err = deleteNpcTask(c, req.NpcIds)
	if err != nil {
		log.Printf("deleteNpcTask failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, nil)
}

func deleteNpcTask(c *gin.Context, ids []string) error {
	npcTaskDb := database.Query.NpcTask
	_, err := npcTaskDb.WithContext(c).Where(npcTaskDb.ID.In(ids...)).Delete(&model.NpcTask{})
	if err != nil {
		log.Printf("npcTaskDb delete failed, err:%v", err)
		return util.BuildErrorInfo("npcTaskDb delete failed, err:%v", err)
	}
	return nil
}
