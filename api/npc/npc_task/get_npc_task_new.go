package npc_task

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

// 获取NPCtask中去重的数据
func GetNPCTaskNewApi(c *gin.Context) {
	var req model_view.GetNpcTaskNewReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ShouldBindJSON GetNpcTaskReq failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	if err := checkGetNPCTaskNewParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	taskList, err := getNpcTaskNew(c, &req)
	if err != nil {
		log.Printf("getNpcTask failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.OperationErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, taskList)
}

func checkGetNPCTaskNewParam(req *model_view.GetNpcTaskNewReq) error {
	if req.PageNo == 0 {
		req.PageNo = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 100
	}
	return nil
}

func getNpcTaskNew(c *gin.Context, req *model_view.GetNpcTaskNewReq) (*model_view.GetNpcTaskResp, error) {
	npcTaskDb := database.Query.NpcTask
	userTaskDb := database.Query.UserTask
	taskDb := database.Query.Task
	userTasks, err := userTaskDb.WithContext(c).Where(userTaskDb.UserID.Eq(req.UserId)).Find()
	if err != nil {
		log.Printf("userTaskDb find failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("userTaskDb find failed, err:%v", err)
	}
	userTaskIdsUnfinished := make([]string, 0)
	userTaskIdsFinished := make([]string, 0)
	for _, item := range userTasks {
		if item.IsFinished == 0 {
			userTaskIdsUnfinished = append(userTaskIdsUnfinished, item.TaskID)
		} else {
			userTaskIdsFinished = append(userTaskIdsFinished, item.TaskID)
		}
	}
	taskLists, err := taskDb.WithContext(c).Where(taskDb.ID.In(userTaskIdsFinished...)).Find()
	if err != nil {
		log.Printf("taskDb find failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("taskDb find failed, err:%v", err)
	}
	// 把不可以重复的放进去
	for _, item := range taskLists {
		if item.CanRepeated == 0 {
			userTaskIdsUnfinished = append(userTaskIdsUnfinished, item.ID)
		}
	}
	sql := npcTaskDb.WithContext(c)
	if req.ID != "" {
		sql = sql.Where(npcTaskDb.ID.Eq(req.ID))
	}
	if req.NpcID != "" {
		sql = sql.Where(npcTaskDb.NpcID.Eq(req.NpcID))
	}
	if len(userTaskIdsUnfinished) > 0 {
		sql = sql.Where(npcTaskDb.TaskID.NotIn(userTaskIdsUnfinished...))
	}
	total, err := sql.Count()
	if err != nil {
		log.Printf("npcTaskDb count failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("npcTaskDb count failed, err:%v", err)
	}
	npcTaskList, err := sql.Offset((req.PageNo - 1) * req.PageSize).Limit(req.PageSize).Find()
	if err != nil {
		log.Printf("npcTaskDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("npcTaskDb query failed, err:%v\n", err)
	}
	taskIds := make([]string, len(npcTaskList))
	for i, item := range npcTaskList {
		taskIds[i] = item.TaskID
	}
	taskList, err := task.GetTaskById(c, taskIds)
	if err != nil {
		log.Printf("task.GetTaskById failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("task.GetTaskById failed, err:%v\n", err)
	}
	taskMap := make(map[string]*model.Task)
	for _, item := range taskList {
		taskMap[item.ID] = item
	}
	npcTaskDetail := make([]*model_view.NpcTaskDetail, len(npcTaskList))
	for i, item := range npcTaskList {
		task, ok := taskMap[item.TaskID]
		if !ok {
			continue
		}
		npcTaskDetail[i] = &model_view.NpcTaskDetail{
			NpcTask: item,
			Task:    task,
		}
	}
	res := &model_view.GetNpcTaskResp{
		Total: total,
		Data:  npcTaskDetail,
	}
	return res, nil
}
