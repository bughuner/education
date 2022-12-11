package npc_task

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

func AddNpcTaskApi(c *gin.Context) {
	var req model.NpcTask
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON npc task failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkAddNpcTaskParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	npcTask, err := addNpcTask(c, req.NpcID, req.TaskID)
	if err != nil {
		log.Printf("addNpcTask failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, npcTask)
}

func checkAddNpcTaskParam(npcTask *model.NpcTask) error {
	if npcTask.TaskID == "" {
		return util.BuildErrorInfo("任务ID为空")
	}
	if npcTask.NpcID == "" {
		return util.BuildErrorInfo("NPC ID为空")
	}
	return nil
}

func addNpcTask(c *gin.Context, npcId, taskId string) (*model.NpcTask, error) {
	taskList, err := task.GetTaskById(c, []string{taskId})
	if err != nil {
		log.Printf("GetTaskById failed, userId:%v, npcId:%v, err:%v\n", npcId, taskId, err)
		return nil, util.BuildErrorInfo("GetTaskById failed,, npcId:%v, taskId:%v, err:%v", npcId, taskId, err)
	}
	if taskList == nil || len(taskList) == 0 {
		log.Printf("task not exist, npcId:%v, taskId:%v, err:%v\n", npcId, taskId, err)
		return nil, util.BuildErrorInfo("task not exist,, npcId:%v, taskId:%v, err:%v", npcId, taskId, err)
	}
	npcTask := &model.NpcTask{
		NpcID:  npcId,
		TaskID: taskId,
	}
	id := util.GetUUID()
	npcTask.ID = id
	npcTaskDb := database.Query.NpcTask
	err = npcTaskDb.WithContext(c).Create(npcTask)
	if err != nil {
		log.Printf("npcTaskDb create failed, npcId:%v, taskId:%v, err:%v", npcId, taskId, err)
		return nil, util.BuildErrorInfo("npcTaskDb create failed, npcId:%v, taskId:%v, err:%v", npcId, taskId, err)
	}
	return npcTask, nil
}
