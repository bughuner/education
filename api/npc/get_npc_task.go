package npc

import (
	"education/api/task"
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func GetNPCTaskApi(c *gin.Context) {
	npcId := c.Query("npc_id")
	if err := checkGetNPCTaskParam(npcId); err != nil {
		log.Printf("checkParam failed, npcId:%v, err:%v\n", npcId, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	taskList, err := getNpcTaskByNpcId(c, npcId)
	if err != nil {
		log.Printf("getNpcTaskByNpcId failed, npcId:%v, err:%v\n", npcId, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, taskList)
}

func checkGetNPCTaskParam(npcId string) error{
	if npcId == "" {
		return util.BuildErrorInfo("参数错误")
	}
	return nil
}

func getNpcTaskByNpcId(c *gin.Context, npcId string) ([]*model.Task, error) {
	npcTaskDb := database.Query.NpcTask
	npcTaskList, err := npcTaskDb.WithContext(c).Where(npcTaskDb.NpcID.Eq(npcId)).Find()
	if err != nil {
		log.Printf("npcTaskDb query failed, npcId:%v, err:%v\n", npcId, err)
		return nil, util.BuildErrorInfo("npcTaskDb query failed, npcId:%v, err:%v\n", npcId, err)
	}
	if len(npcTaskList) == 0 {
		return nil, nil
	}
	taskIds := make([]string, len(npcTaskList))
	for i, item := range npcTaskList {
		taskIds[i] = item.TaskID
	}
	taskList, err := task.GetTaskById(c, taskIds)
	if err != nil {
		log.Printf("GetTaskById failed, npcId:%v, err:%v",  npcId, err)
		return nil, util.BuildErrorInfo("GetTaskById failed, npcId:%v, err:%v",  npcId, err)
	}
	return taskList, nil
}