package task

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func AddTaskApi(c *gin.Context) {
	var req model.Task
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON task failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkAddParam(&req); err != nil {
		log.Printf("checkAddParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	task, err := addTask(c, &req)
	if err != nil {
		log.Printf("addTaskList failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, task)
}

func checkAddParam(req *model.Task) error {
	if req.Type == 0 {
		return util.BuildErrorInfo("类型不能为0")
	}
	if req.Level == 0 {
		return util.BuildErrorInfo("等级不能为0")
	}
	if req.Coin == 0 {
		return util.BuildErrorInfo("金币不能为0")
	}
	if req.Experience == 0 {
		return util.BuildErrorInfo("经验值不能为0")
	}
	if req.Introduction == "" {
		return util.BuildErrorInfo("简介不能为0")
	}
	return nil
}

func addTask(c *gin.Context, task *model.Task) (*model.Task, error) {
	taskDb := database.Query.Task
	id := util.GetUUID()
	taskEntity := &model.Task{
		ID:           id,
		Type:         task.Type,
		Level:        task.Level,
		Introduction: task.Introduction,
		Image:        task.Image,
		Experience:   task.Experience,
		Coin:         task.Coin,
		Num:          task.Num,
		PreTask:      task.PreTask,
		TargetID:     task.TargetID,
	}
	err := taskDb.WithContext(c).Save(taskEntity)
	if err != nil {
		log.Printf("taskDb save failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("taskDb save failed, err:%v\n", err)
	}
	return taskEntity, nil
}
