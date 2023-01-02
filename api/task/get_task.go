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

func GetTaskApi(c *gin.Context) {
	var req model_view.TaskReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON model_view.TaskReq failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err = checkGetParam(&req); err != nil {
		log.Printf("checkGetParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	taskList, err := getTaskList(c, &req)
	if err != nil {
		log.Printf("getTaskList failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.InternalServerError, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, taskList)
}

func checkGetParam(req *model_view.TaskReq) error {
	if req.PageNo == 0 {
		req.PageNo = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 100
	}
	return nil
}

func getTaskList(c *gin.Context, req *model_view.TaskReq) (*model_view.TaskResp, error) {
	taskDb := database.Query.Task
	sql := taskDb.WithContext(c)
	if req.ID != "" {
		sql = sql.Where(taskDb.ID.Eq(req.ID))
	}
	if req.Type != 0 {
		sql = sql.Where(taskDb.Level.Eq(req.Type))
	}
	if req.Level != 0 {
		sql = sql.Where(taskDb.Level.Eq(req.Level))
	}
	if req.PreTask != "" {
		sql = sql.Where(taskDb.PreTask.Eq(req.PreTask))
	}
	if req.Num != 0 {
		sql = sql.Where(taskDb.Num.Eq(req.Num))
	}
	if req.Coin != 0 {
		sql = sql.Where(taskDb.Coin.Eq(req.Coin))
	}
	if req.Experience != 0 {
		sql = sql.Where(taskDb.Experience.Eq(req.Experience))
	}
	total, err := sql.Count()
	if err != nil {
		log.Printf("taskDb count failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("taskDb count failed, err:%v", err)
	}
	taskList, err := sql.Offset((req.PageNo - 1) * req.PageSize).Limit(req.PageSize).Find()
	if err != nil {
		log.Printf("taskDb count query, err:%v\n", err)
		return nil, util.BuildErrorInfo("taskDb query failed, err:%v", err)
	}
	res := &model_view.TaskResp{
		Total: total,
		Data:  taskList,
	}
	return res, nil
}

func GetTaskById(c *gin.Context, taskIds []string) ([]*model.Task, error) {
	taskDb := database.Query.Task
	taskList, err := taskDb.WithContext(c).Where(taskDb.ID.In(taskIds...)).Find()
	if err != nil {
		log.Printf("taskDb query failed, taskIds:%v, err:%v\n", taskIds, err)
		return nil, util.BuildErrorInfo("taskDb query failed, taskIds:%v, err:%v", taskIds, err)
	}
	return taskList, nil
}
