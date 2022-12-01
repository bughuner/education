package user

import (
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/util"
	"github.com/gin-gonic/gin"
	"log"
)

func AddUserTaskApi(c *gin.Context) {
	var req model.UserTask
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON user task failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkAddUserTaskParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	userTask, err := addUserTask(c, req.UserID, req.TaskID)
	if err != nil {
		log.Printf("addUserTask failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, userTask)
}

func checkAddUserTaskParam(userTask *model.UserTask) error {
	if userTask.UserID == "" {
		return util.BuildErrorInfo("用户ID为空")
	}
	if userTask.TaskID == "" {
		return util.BuildErrorInfo("任务ID为空")
	}
	return nil
}

func addUserTask(c *gin.Context, userId, taskId string) (*model.UserTask, error) {
	userTask := &model.UserTask{
		UserID: userId,
		TaskID: taskId,
	}
	id := util.GetUUID()
	userTask.ID = id
	userTaskDb := database.Query.UserTask
	err := userTaskDb.WithContext(c).Create(userTask)
	if err != nil {
		log.Printf("userTaskDb create failed, userId:%v, taskId:%v, err:%v", userId, taskId, err)
		return nil, util.BuildErrorInfo("userTaskDb create failed, userId:%v, taskId:%v, err:%v", userId, taskId, err)
	}
	return userTask, nil
}