package user_task

import (
	"education/api/task"
	"education/common"
	errno "education/common/erron"
	"education/consts"
	"education/database"
	"education/model"
	"education/service"
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
	taskList, err := task.GetTaskById(c, []string{taskId})
	if err != nil {
		log.Printf("GetTaskById failed, userId:%v, taskId:%v, err:%v\n", userId, taskId, err)
		return nil, util.BuildErrorInfo("GetTaskById failed,, userId:%v, taskId:%v, err:%v", userId, taskId, err)
	}
	if len(taskList) == 0 {
		log.Printf("task not exist, userId:%v, taskId:%v, err:%v\n", userId, taskId, err)
		return nil, util.BuildErrorInfo("task not exist,, userId:%v, taskId:%v, err:%v", userId, taskId, err)
	}
	taskEntity := taskList[0]
	userTask := &model.UserTask{
		UserID:     userId,
		TaskID:     taskId,
		Type:       taskEntity.Type,
		IsFinished: 0,
		Count:      taskEntity.Num,
	}
	if taskEntity.Type == consts.TaskTypeDoc {
		docList, err := service.GetDocQuestionByDocId(c, []string{taskEntity.TargetID})
		if err != nil {
			log.Printf("service GetDocQuestionByDocId failed, err:%v\n", err)
			return nil, util.BuildErrorInfo("service GetDocQuestionByDocId failed, err:%v\n", err)
		}
		if len(docList) == 0 {
			return nil, util.BuildErrorInfo("关联的文章ID有错误")
		}
		userTask.Count = int64(len(docList))
	}

	id := util.GetUUID()
	userTask.ID = id
	userTaskDb := database.Query.UserTask
	err = userTaskDb.WithContext(c).Create(userTask)
	if err != nil {
		log.Printf("userTaskDb create failed, userId:%v, taskId:%v, err:%v", userId, taskId, err)
		return nil, util.BuildErrorInfo("userTaskDb create failed, userId:%v, taskId:%v, err:%v", userId, taskId, err)
	}
	return userTask, nil
}
