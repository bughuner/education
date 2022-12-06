package task

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

func UpdateTaskApi(c *gin.Context) {
	var req model.Task
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON task failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkUpdateParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	task, err := updateTask(c, &req)
	if err != nil {
		log.Printf("updateMonster failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, task)
}

func checkUpdateParam(req *model.Task) error {
	if req.ID == "" {
		return util.BuildErrorInfo("ID为空")
	}
	return nil
}

func updateTask(c *gin.Context, req *model.Task) (*model.Task, error) {
	taskDb := database.Query.Task
	oldTask, err := taskDb.WithContext(c).Where(taskDb.ID.Eq(req.ID)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, util.BuildErrorInfo("taskDb query failed, err:%v", err)
	}
	if oldTask == nil {
		return nil, util.BuildErrorInfo("任务不存在")
	}
	if req.Type != 0 {
		oldTask.Type = req.Type
	}
	if req.Coin != 0 {
		oldTask.Coin = req.Coin
	}
	if req.PreTask != "" {
		oldTask.PreTask = req.PreTask
	}
	if req.Num != 0 {
		oldTask.Num = req.Num
	}
	if req.Experience != 0 {
		oldTask.Experience = req.Experience
	}
	if req.Image != "" {
		oldTask.Image = req.Image
	}
	if req.Level != 0 {
		oldTask.Level = req.Level
	}
	err = taskDb.WithContext(c).Save(oldTask)
	if err != nil {
		log.Printf("taskDb save failed, req:%v, err:%v", req, err)
		return nil, util.BuildErrorInfo("monsterDb save failed, err:%v", err)
	}
	return oldTask, nil
}
