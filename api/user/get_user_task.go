package user

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

func GetUserTaskApi(c *gin.Context) {
	userId := c.Query("user_id")
	if err := checkGetUserTaskParam(userId); err != nil {
		log.Printf("checkParam failed, user_id:%v, err:%v\n", userId, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	taskList, err := getUserTaskByUserId(c, userId)
	if err != nil {
		log.Printf("getUserTaskByUserId failed, user_id:%v, err:%v\n", userId, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, taskList)
}

func checkGetUserTaskParam(id string) error{
	if id == "" {
		return util.BuildErrorInfo("参数错误")
	}
	return nil
}

func getUserTaskByUserId(c *gin.Context, userId string) ([]*model.Task, error) {
	userTaskDb := database.Query.UserTask
	userTaskList, err := userTaskDb.WithContext(c).Where(userTaskDb.UserID.Eq(userId)).Find()
	if err != nil {
		log.Printf("userTaskDb query failed, userId:%v, err:%v\n", userId, err)
		return nil, util.BuildErrorInfo("userTaskDb query failed, userId:%v, err:%v\n", userId, err)
	}
	if len(userTaskList) == 0 {
		return nil, nil
	}
	taskIds := make([]string, len(userTaskList))
	for i, item := range userTaskList {
		taskIds[i] = item.TaskID
	}
	taskList, err := task.GetTaskById(c, taskIds)
	if err != nil {
		log.Printf("GetTaskById failed, userId:%v, err:%v",  userId, err)
		return nil, util.BuildErrorInfo("GetTaskById failed, userId:%v, err:%v",  userId, err)
	}
	return taskList, nil
}
