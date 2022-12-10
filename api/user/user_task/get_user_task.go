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
	var req model_view.GetUserTaskReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON GetUserTaskReq failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err := checkGetUserTaskParams(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	taskList, err := getUserTask(c, &req)
	if err != nil {
		log.Printf("getUserTask failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, taskList)
}

func checkGetUserTaskParams(req *model_view.GetUserTaskReq) error {
	if req.PageNo == 0 {
		req.PageNo = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 100
	}
	return nil
}

func getUserTask(c *gin.Context, req *model_view.GetUserTaskReq) (*model_view.GetUserTaskResp, error) {
	userTaskDb := database.Query.UserTask
	sql := userTaskDb.WithContext(c)
	if req.ID != "" {
		sql = sql.Where(userTaskDb.ID.Eq(req.ID))
	}
	if req.UserID != "" {
		sql = sql.Where(userTaskDb.UserID.Eq(req.UserID))
	}
	total, err := sql.Count()
	if err != nil {
		log.Printf("userTaskDb count failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("userTaskDb query failed, err:%v", err)
	}
	userTaskList, err := sql.Offset(req.PageNo - 1).Limit(req.PageSize).Find()
	if err != nil {
		log.Printf("userTaskDb count failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("userTaskDb query failed, err:%v", err)
	}
	if userTaskList == nil || len(userTaskList) == 0 {
		return nil, nil
	}
	taskIds := make([]string, len(userTaskList))
	for i, item := range userTaskList {
		taskIds[i] = item.TaskID
	}
	taskList, err := task.GetTaskById(c, taskIds)
	if err != nil {
		log.Printf("GetTaskById failed, taskIds:%v, err:%v\n", taskIds, err)
		return nil, util.BuildErrorInfo("GetTaskById failed, taskIds:%v, err:%v", taskIds, err)
	}
	m := make(map[string]*model.Task)
	for _, item := range taskList {
		m[item.ID] = item
	}
	userTasks := make([]*model_view.UserTaskResp, len(userTaskList))
	for i, item := range userTaskList {
		userTasks[i] = &model_view.UserTaskResp{
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
		userTasks[i].Level = task.Level
		userTasks[i].Introduction = task.Introduction
		userTasks[i].Image = task.Image
		userTasks[i].Experience = task.Experience
		userTasks[i].Count = task.Coin
		userTasks[i].Num = task.Num
		userTasks[i].PreTask = task.PreTask
	}
	res := &model_view.GetUserTaskResp{
		Total: total,
		Data:  userTasks,
	}
	return res, nil
}
