package npc_task

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

func UpdateNpcTaskApi(c *gin.Context) {
	var req model.NpcTask
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON NpcTask failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkUpdateParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	npcTask, err := updateNpcTask(c, &req)
	if err != nil {
		log.Printf("updateNpcTask failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, npcTask)
}

func checkUpdateParam(plat *model.NpcTask) error {
	if plat.ID == "" {
		return util.BuildErrorInfo("ID为空")
	}
	return nil
}

func updateNpcTask(c *gin.Context, req *model.NpcTask) (*model.NpcTask, error) {
	npcTaskDb := database.Query.NpcTask
	npcTask, err := npcTaskDb.WithContext(c).Where(npcTaskDb.ID.Eq(req.ID)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("npcTaskDb query failed, id:%v, err:%v\n", req.ID, err)
		return nil, util.BuildErrorInfo("npcTaskDb query failed, err:%v", err)
	}
	if npcTask == nil {
		return nil, util.BuildErrorInfo("npc任务不存在")
	}
	if req.NpcID != "" {
		npcTask.NpcID = req.NpcID
	}
	if req.TaskID != "" {
		npcTask.TaskID = req.TaskID
	}
	err = npcTaskDb.WithContext(c).Save(npcTask)
	if err != nil {
		log.Printf("npcTaskDb save failed, npcTask:%v, err:%v", npcTask, err)
		return nil, util.BuildErrorInfo("npcTaskDb save failed, err:%v", err)
	}
	return npcTask, nil
}
