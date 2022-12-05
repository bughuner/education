package user_task

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

func UpdateUserTaskApi(c *gin.Context) {
	var req model.UserTask
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON user_task failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err := checkUpdateUseTaskParams(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	userQuestion, err := updateUserTask(c, req.ID, req.UserID, req.TaskID, req.Count)
	if err != nil {
		log.Printf("updateUserTask failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, userQuestion)
}

func checkUpdateUseTaskParams(req *model.UserTask) error {
	if req.ID == "" {
		return util.BuildErrorInfo("ID为空")
	}
	if req.UserID == "" {
		return util.BuildErrorInfo("用户ID为空")
	}
	if req.TaskID == "" {
		return util.BuildErrorInfo("任务ID为空")
	}
	return nil
}

func updateUserTask(c *gin.Context, id, userId, taskId string, count int64) (*model.UserTask, error) {
	userTaskDb := database.Query.UserTask
	userTask, err := userTaskDb.WithContext(c).Where(userTaskDb.ID.Eq(id)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("userTaskDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("userTaskDb query failed, err:%v", err)
	}
	if userTask == nil {
		return nil, util.BuildErrorInfo("用户任务不存在")
	}
	if userId != "" {
		userTask.UserID = userId
	}
	if taskId != "" {
		userTask.TaskID = taskId
	}
	userTask.Count = count
	if userTask.Count <= 0 {
		userTask.IsFinished = 1
	}
	err = userTaskDb.WithContext(c).Save(userTask)
	if err != nil {
		log.Printf("userQuestionDb save failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("userQuestionDb save failed, err:%v", err)
	}
	return userTask, nil
}
