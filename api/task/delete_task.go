package task

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

func DeleteTaskApi(c *gin.Context) {
	var req model_view.DeleteTaskReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON task_ids failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	err = deleteTasks(c, req.TaskIds)
	if err != nil {
		log.Printf("getTaskList failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, nil)
}

func deleteTasks(c *gin.Context, taskIds []string) error {
	taskDb := database.Query.Task
	_, err := taskDb.WithContext(c).Where(taskDb.ID.In(taskIds...)).Delete(&model.Task{})
	if err != nil {
		log.Printf("taskDb dekete failed, err:%v\n", err)
		return util.BuildErrorInfo("taskDb dekete failed, err:%v", err)
	}
	return nil
}
