package user_task

import (
	"education/api/task"
	"education/common"
	errno "education/common/erron"
	"education/database"
	"education/model"
	"education/model/model_view"
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

func checkGetUserTaskParam(id string) error {
	if id == "" {
		return util.BuildErrorInfo("参数错误")
	}
	return nil
}

func getUserTaskByUserId(c *gin.Context, userId string) ([]*model_view.UserTaskResp, error) {
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
		log.Printf("GetTaskById failed, userId:%v, err:%v", userId, err)
		return nil, util.BuildErrorInfo("GetTaskById failed, userId:%v, err:%v", userId, err)
	}
	m := make(map[string]*model.Task)
	for _, item := range taskList {
		m[item.ID] = item
	}
	res := make([]*model_view.UserTaskResp, len(userTaskList))
	for i, item := range userTaskList {
		res[i] = &model_view.UserTaskResp{
			ID:         item.ID,
			UserID:     item.UserID,
			TaskID:     item.TaskID,
			Type:       item.Type,
			IsFinished: item.IsFinished,
			Count:      item.Count,
		}
		task, ok := m[item.TaskID]
		if !ok {
			log.Printf("task not found")
			continue
		}
		res[i].Level = task.Level
		res[i].Introduction = task.Introduction
		res[i].Image = task.Image
		res[i].Experience = task.Experience
		res[i].Count = task.Coin
		res[i].Num = task.Num
		res[i].PreTask = task.PreTask
	}
	return res, nil
}
